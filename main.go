package main

import (
	"fmt"

	ora "gopkg.in/rana/ora.v4"
)

func main() {
	fmt.Println("Start..")

	// call stored procedure
	// pass *Rset to Exe to receive the results of a sys_refcursor
	stmtProcCall, err := ses.Prep("CALL SP_DEMOG(:1)")
	defer stmtProcCall.Close()
	if err != nil {
		panic(err)
	}
	procRset := &ora.Rset{}
	rowsAffected, err = stmtProcCall.Exe(procRset)
	if err != nil {
		panic(err)
	}
	if procRset.IsOpen() {
		for procRset.Next() {
			fmt.Println(procRset.Row[0], procRset.Row[1])
		}
		if err := procRset.Err(); err != nil {
			panic(err)
		}
		fmt.Println(procRset.Len())
	}

	fmt.Println("End..")
}
