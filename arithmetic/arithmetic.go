package arithmetic

import (
	"errors"
	"fmt"
)

func Add(a int, b int) (int, error) {

	result := a + b
	message := fmt.Sprintf("The sum of %v and %v is %v.", a, b, result)
	fmt.Println(message)

	return result, nil
}

func Sub(a int, b int) (int, error) {

	result := a - b
	message := fmt.Sprintf("The difference of %v and %v is %v.", a, b, result)
	fmt.Println(message)

	return result, nil
}

func Mul(a int, b int) (int, error) {

	result := a * b
	message := fmt.Sprintf("The multiplication of %v and %v is %v.", a, b, result)
	fmt.Println(message)

	return result, nil
}

func Div(a int, b int) (int, error) {

	if b == 0 {
		error_message := fmt.Sprintf("The division of %v and %v is impossible.", a, b)
        return 0, errors.New(error_message)
    }

	result := a / b
	message := fmt.Sprintf("The division of %v and %v is %v.", a, b, result)
	fmt.Println(message)

	return result, nil
}
