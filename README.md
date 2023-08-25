A set of GoLang programming challenges TUI app.

## Installation and usage

Copy and paste into terminal. You should have Go installed.

### Install

```
go install github.com/rusinikita/trainer@latest
```

### Run

```
trainer
```

## UI example

```
                                                                                                                      
   Structs                                                                                                            
                                                                                                                      
  >  0. type Employee struct {                                                                                        
     1.     AccessLevel string                                                                                        
     2. }                                                                                                             
     3.                                                                                                               
     4. func upgradeAccess(employee *Emp                                                                              
     5.     employee = &Employee{                                                                                     
     6.         AccessLevel: "Admin",                                                                                 
     7.     }                                                                                                         
     8. }                                                                                                             
     9.                                                                                                               
    10. func main() {                                                                                                 
    11.     employee := &Employee{                                                                                    
    12.         AccessLevel: "User",                                                                                  
    13.     }                                                                                                         
    14.     fmt.Println(employee.AccessL                                                                              
    15.     upgradeAccess(employee)                                                                                   
    16.     fmt.Println(employee.AccessL                                                                              
    17. }                                                                                                             
                                                                                                                      
                                                                                                                      
──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
                                                                                                                      
 Question 0: Select program output                                                                                    
                                                                                                                      
╔═══════╗┌───────┐┌──────┐┌───────┐                                                                                   
║ User  ║│ Admin ││ User ││ Admin │                                                                                   
║ Admin ║│ User  ││ User ││ Admin │                                                                                   
╚═══════╝└───────┘└──────┘└───────┘                                                                                   
                                                                                                                      
  ↑/k up • ↓/j down • ← left • → right • ⮐ select • q quit • ? more                                                   
                                                                                                                      
```