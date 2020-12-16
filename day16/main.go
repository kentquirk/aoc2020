package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/kentquirk/stringset/v2"
)

// ValidRange has a min and max and can check for containment
// of an int within the range.
type ValidRange struct {
	min int
	max int
}

// In returns true if the argument is in range.
func (r ValidRange) In(v int) bool {
	return r.min <= v && v <= r.max
}

// Field defines a field on a ticket
type Field struct {
	name   string
	valids []ValidRange
}

// Valid returns true if the value is valid for this field.
func (f *Field) Valid(v int) bool {
	for _, r := range f.valids {
		if r.In(v) {
			return true
		}
	}
	return false
}

// TicketField is one field on a ticket.
// It starts without a name, but we may add one later.
type TicketField struct {
	nameset *stringset.StringSet
	value   int
}

// NewTicketField constructs an unnamed ticketfield
func NewTicketField(s string) []TicketField {
	fields := make([]TicketField, 0)
	for _, v := range strings.Split(s, ",") {
		n, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		fields = append(fields, TicketField{
			nameset: stringset.New(),
			value:   n,
		},
		)
	}
	return fields
}

// NewField creates and populates a Field object.
func NewField(s string) *Field {
	pat := regexp.MustCompile("([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)")
	groups := pat.FindStringSubmatch(s)
	if groups == nil {
		return nil
	}
	a1, _ := strconv.Atoi(groups[2])
	a2, _ := strconv.Atoi(groups[3])
	b1, _ := strconv.Atoi(groups[4])
	b2, _ := strconv.Atoi(groups[5])
	return &Field{
		name: groups[1],
		valids: []ValidRange{
			{
				min: a1,
				max: a2,
			},
			{
				min: b1,
				max: b2,
			},
		},
	}
}

func parseFields(s string) []*Field {
	fields := make([]*Field, 0)
	for _, line := range strings.Split(s, "\n") {
		if f := NewField(line); f != nil {
			fields = append(fields, f)
		}
	}
	return fields
}

func parseTix(s string) [][]TicketField {
	tix := make([][]TicketField, 0)
	for _, line := range strings.Split(s, "\n") {
		if fields := NewTicketField(line); len(fields) != 0 {
			tix = append(tix, fields)
		}
	}
	return tix
}

func accumulateInvalids(fields []*Field, tix [][]TicketField) int {
	total := 0
	for _, t := range tix {
		for _, v := range t {
			isValid := false
			for _, f := range fields {
				if f.Valid(v.value) {
					isValid = true
					break
				}
			}
			if !isValid {
				total += v.value
				continue
			}
		}
	}
	return total
}

func discardInvalids(fields []*Field, tix [][]TicketField) [][]TicketField {
	validTix := make([][]TicketField, 0)
	for _, t := range tix {
		tixValid := true
		for _, v := range t {
			valueValid := false
			for _, f := range fields {
				if f.Valid(v.value) {
					valueValid = true
					break
				}
			}
			if !valueValid {
				tixValid = false
				break
			}
		}
		if tixValid {
			validTix = append(validTix, t)
		}
	}
	return validTix
}

func assignFields(fields []*Field, tix [][]TicketField) map[string]int {
	for tixIx := range tix {
		for _, f := range fields {
			for valueIx := range tix[tixIx] {
				value := tix[tixIx][valueIx].value
				if f.Valid(value) {
					tix[tixIx][valueIx].nameset.Add(f.name)
				}
			}
		}
	}
	// dumpTix(tix[:2])

	// we need to build a []stringset from the possible fieldnames
	// for each entry
	possibles := make([]*stringset.StringSet, 0)

	// now go through all the fields in the first ticket and
	// create the possibles
	for fieldIx := range tix[0] {
		intset := tix[0][fieldIx].nameset
		for tixIx := range tix {
			intset = intset.Intersection(tix[tixIx][fieldIx].nameset)
		}
		possibles = append(possibles, intset)
	}

	// for _, p := range possibles {
	// 	fmt.Printf("poss: %s\n", p.Join(","))
	// }

	// here we loop through the possibles list looking for
	// name sets that contain only one value; we emit those
	// and then remove that item from the rest of them
	// and repeat until we have eliminated all the possibilities.
	fieldnames := make(map[string]int)
	removes := stringset.New()
	done := false
	count := 0
	for !done {
		done = true // until shown to be false
		for fieldIx := range possibles {
			poss := possibles[fieldIx]
			// remove all the ones we've already done
			poss = poss.Difference(removes)
			// if the result is of length 1, we can record it and
			// add it to the removes list
			switch poss.Length() {
			case 0:
			case 1:
				fieldnames[poss.Join("")] = fieldIx
				removes.Add(poss.Strings()...)
			default:
				done = false
			}
			possibles[fieldIx] = poss
		}
		count++
		if count > 40 {
			// for _, p := range possibles {
			// 	fmt.Println(p.Join(" "))
			// }
			break
		}
	}
	return fieldnames
}

func dumpTix(tix [][]TicketField) {
	for tixIx := range tix {
		for valueIx := range tix[tixIx] {
			fmt.Printf("%s %d   \n", tix[tixIx][valueIx].nameset.Join(","), tix[tixIx][valueIx].value)
		}
		fmt.Println("---")
	}
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	pat := regexp.MustCompile("tickets?:")
	filegroups := pat.Split(string(b), -1)
	fields := parseFields(filegroups[0])
	myTix := parseTix(filegroups[1])[0]
	otherTix := parseTix(filegroups[2])

	fmt.Println(accumulateInvalids(fields, otherTix))
	validTix := discardInvalids(fields, otherTix)
	// dumpTix(validTix)
	fieldnames := assignFields(fields, validTix)
	product := 1
	for f, ix := range fieldnames {
		if strings.HasPrefix(f, "departure") {
			product *= myTix[ix].value
		}
	}
	fmt.Println(product)
}
