package stats

import (
	"github.com/OAbdulloev87/bank/pkg/bank/types"
	
)
// Avg рассчитывает среднюю сумму платежа.
func Avg(payments []types.Payment) types.Money { 
	 sum := types.Money(0) 
	for _, payment := range payments {  
		 sum+= payment.Amount 
	 } 
	 return sum/types.Money(len(payments))
	 }