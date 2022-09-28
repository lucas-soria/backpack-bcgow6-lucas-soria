package ejercicio4

import (
	"errors"
	"fmt"
)

const (
	ErrorHours     = "the employee cannot work less than 80hs in a monthly bases"
	ErrorNegative  = "negative values are not accepted"
	MinimumTaxable = 150_000
	Tax            = 0.1
)

/*
Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
Desarrolla las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de
impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver
un error. El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un
número negativo.

Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”,
“fmt.Errorf()” y “errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función
“main()”.
*/

type Employee struct {
	hoursWorkedMonth []float64
	hourValue        float64
}

func (employee Employee) checkHoursWorkedMonth() (err error) {
	for _, hoursWorked := range employee.hoursWorkedMonth {
		switch {
		case hoursWorked < 0:
			err = errors.New(ErrorNegative)
			fallthrough
		case hoursWorked < 80:
			if err != nil {
				return fmt.Errorf("multiple errors while checking hours worked per month:\n\t- %s\n\t- %w", ErrorHours, err)
			}
			return fmt.Errorf("error while checking hours worked per month:\n\t- %s", ErrorHours)
		}
	}
	return
}

func (employee Employee) checkHourValue() (err error) {
	if employee.hourValue < 0 {
		return fmt.Errorf("error while checking hour's value:\n\t- %s", ErrorNegative)
	}
	return
}

func (employee Employee) checkErrors() (errs error) {
	if hoursErr := employee.checkHoursWorkedMonth(); hoursErr != nil {
		errs = hoursErr
	}
	if valueErr := employee.checkHourValue(); valueErr != nil {
		if errs != nil {
			return fmt.Errorf("%w\n%w", errs, valueErr)
		}
		return valueErr
	}
	return
}

func (employee Employee) maxSalary() (max float64) {
	max = employee.hoursWorkedMonth[0] * employee.hourValue
	for _, hours := range employee.hoursWorkedMonth {
		newMax := hours * employee.hourValue
		if newMax > max {
			max = newMax
		}
	}
	return
}

func (employee Employee) lastMonthSalary() (salary float64, err error) {
	if err := employee.checkErrors(); err != nil {
		return 0, err
	}
	salary = employee.hourValue * employee.hoursWorkedMonth[len(employee.hoursWorkedMonth)-1]
	if salary > MinimumTaxable {
		salary *= 1 - Tax
	}
	return
}

func (employee Employee) halfBonus() (halfBonus float64, err error) {
	if err := employee.checkErrors(); err != nil {
		return 0, err
	}
	halfBonus = employee.maxSalary() / 12 * float64(len(employee.hoursWorkedMonth)) / 2
	return
}

func SalaryTool() {
	employee := Employee{
		hoursWorkedMonth: []float64{85, -129, 118, 117},
		hourValue:        8,
	}
	fmt.Println("Last logged moth's salary...")
	lastMonthSalary, err := employee.lastMonthSalary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Last logged moth's salary is:", lastMonthSalary)
	}
	fmt.Println("Half of the bonus...")
	halfBonus, err := employee.halfBonus()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Half of the bonus:", halfBonus)
	}
}
