package string_sum

import (
	"fmt"
	"errors"
	"strconv"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	in:=[]rune(input)
	var r []int64
	t:=""
	var s int64 
	for i:=0;i<len(in);i++ {
		if string(in[i])==" " {

		} else if string(in[i])=="-" || string(in[i])=="+" {
			if t!="" {
				n,_:=strconv.ParseInt(t,10,32)
				r=append(r,n)
				t=string(in[i])
			} else if string(in[i])=="-" {
				t=t+"-"
			}
		} else{
			_,err:=strconv.ParseInt(string(in[i]),10,32)
			if err != nil {
				return "",fmt.Errorf("Input expression is not valid: "+string(in[i])+" %w",err)
			}
			t=t+string(in[i])
		}
	}
	if t!="" {
		n,_:=strconv.ParseInt(t,10,32)
		r=append(r,n)
	}
	if(len(r)==0){
		return "",fmt.Errorf("Input expression is not valid: %w",errorEmptyInput)
	}
	if(len(r)!=2){
		return "",fmt.Errorf("Input expression is not valid: %w",errorNotTwoOperands)
	}
	for i:=0;i<len(r);i++ {
		s+=r[i]
	}
	return strconv.FormatInt(s, 10), nil
}
