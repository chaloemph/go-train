package main

import (
	"fmt"
	s "strings"
)

func main() {
	var p = fmt.Println
	name := "amang nima"
	job := " web developer"
	fmt.Println(name)
	fmt.Println(name[0:5])
	fmt.Println(name[1:])
	fmt.Println(name[:5])

	fmt.Println(name, job)
	fmt.Println(name + job)

	p("Contains : ", s.Contains(job, "web"))
	p("Count : ", s.Count(job, "web"))
	p("HasPrefix : ", s.HasPrefix(job, " web"))
	p("HasSuffix : ", s.HasSuffix(job, "per"))
	p("Index : ", s.Index(job, "developer"))
	p("Repeat : ", s.Repeat(job, 5))
	p("Replace : ", s.Replace(job, "e", "o", -1)) //replace all when match
	p("Replace : ", s.Replace(job, "e", "o", 1))  //replace 1 when match
	p("Split : ", s.Split(job, "e"))
	p("ToLower : ", s.ToLower(job))
	p("ToUpper : ", s.ToUpper(job))
	p("Len : ", len(job))
}
