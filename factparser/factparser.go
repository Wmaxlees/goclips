package factparser

import ("strings"
        "strconv")

type FactList struct {
    vals map[int]string
}

func NewFactParser(facts string) FactList {
    var result FactList

    lines := strings.Split(facts, "\n")

    for _, element := range lines {
        pieces := strings.Split(element, "     ")
        indexString := strings.TrimPrefix(pieces[0], "f-")

        index, _ := strconv.Atoi(indexString)

        result.vals[index] = pieces[1]
    }

    return result
}

func (this* FactList) GetFact(id int) string {
    return this.vals[id]
}