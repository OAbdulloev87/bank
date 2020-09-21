package stats

import (
	"fmt"
	"github.com/OAbdulloev87/bank/pkg/bank/types"
	)
	func ExampleAvg(){
		payments := []types.Payment{
		  {
			ID:2,
			Amount:43_00,
			Category: "Cat",
		  },
		  {
			ID:1,
			Amount:41_00,
			Category: "Cat",
		  },
		  {
			ID:3,
			Amount:42_00,
			Category: "Cat",
		  },
		}
	  
		fmt.Println(Avg(payments))
	  
		//Output: 4200
	  }
	  