package main

import "fmt"

const (
	banner = `
	
	___                      _               
	|       |    \ /  |                       
	 -+-    +-    +   |       +   |-     -    
		|   |\   /    |       |   | |   |/    
	 ---               ---               --   
											  
	
	`

	banner2 = `
	                          ▆
	___                                     
	┃       ┃    ╲ ╱  ┃       ┃	                
	╰─╮     ┃━    ╋   ┃       ╋   ┏━┓   ┃━━━	  
	  ┃     ┃╲   ╱    ┃       ┃━━━┃ ┃   ┃━━━    
    	┈┈┛               ┗━━━━━━━┛     ┃━━━┃━━━
	`

	logo = `
	 ┏━┓
	┃┃ ┃
    ━━━━┛
	Sky Line
	* *
	 * * 
	  * 
	* *    
	**     
	*   
		**
		
		
	*    
		*  
	`

	stars = `
					* 	 *
					*	  c* 
						* 
					* 		*    
					*	*     
							*   
					*
						*
					*    
						* 
	`
)

func main() {
	fmt.Println("\n\n\n\n")
	fmt.Println("\t\t\t\t	 \033[38;5;51m┏━┓\x1b[0m")
	fmt.Println("\t\t\t\t	\033[38;5;56m┃\033[38;5;51m┃ ┃\x1b[0m")
	fmt.Println("\t\t\t\t    \033[38;5;56m━━━━┛\x1b[0m")
	fmt.Println("\t\t\t\t	\033[38;5;249mSky Line")
	fmt.Println(stars)
}
