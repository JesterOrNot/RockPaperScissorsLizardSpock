package main
import "fmt"
import "math/rand"
import "time"
func cpu() string {
	var decision string
	rand.Seed(time.Now().UnixNano())
	var cpuChoice int = rand.Intn(5) + 1
	switch cpuChoice {
	case 1:
		decision = "rock"
	case 2:
		decision = "paper"
	case 3:
		decision = "scissors"
    case 4:
        decision = "spock"
    case 5:
        decision = "lizard"
    }
	return decision
}
func user() string {
	var decision string
	var userChoice string
	fmt.Println("Welcome to rock paper scissors lizard spock!")
	fmt.Print("Your options are rock, paper, scissors, lizard, and spock what are you?: ")
	_, err := fmt.Scan(&userChoice)
	if userChoice == "rock" {
		decision = "rock"
	} else if userChoice == "paper" {
		decision = "paper"
	} else if userChoice == "scissors" {
		decision = "scissors"
	} else if userChoice == "lizard" {
        decision = "lizard"
    } else if userChoice == "spock" {
        decision = "spock"
    }
	_ = err
	return decision
}
func continueOrNo() {
	var awnser string
	fmt.Print("Do you want to go again?(y/n): ")
	fmt.Scan(&awnser)
	if awnser == "y" {
		main()
	} else if awnser == "n" {
		fmt.Println("Ok cool, Have a Good day!")
	} else {
		fmt.Println("Sorry, response invalid! Please input (y/n)")
		continueOrNo()
	}
}
func main() {
	user := user()
	cpu := cpu()
	if cpu == user {
		fmt.Println("Draw both users picked", cpu)
		continueOrNo()
	} else if cpu == "paper" && user == "rock" {
		fmt.Println("Paper covers rock! You lose!")
		continueOrNo()
	} else if cpu == "paper" && user == "scissors" {
		fmt.Println("Scissors cuts paper! You win!")
		continueOrNo()
	} else if cpu == "paper" && user == "lizard" {
        fmt.Println("Lizard eats paper! You Win!")
        continueOrNo()
    } else if cpu == "paper" && user == "spock" {
        fmt.Println("Paper disproves spock! You lose!")
        continueOrNo()
    } else if cpu == "rock" && user == "paper" {
		fmt.Println("Paper covers rock! You win!")
		continueOrNo()
	} else if cpu == "rock" && user == "scissors" {
		fmt.Println("Rock crushes Scissors! You lose!")
		continueOrNo()
	} else if cpu == "rock" && user == "lizard"  {
        fmt.Println("Rock crushes lizard! You lose!")
        continueOrNo()
    } else if cpu == "rock" && user == "spock" {
        fmt.Println("Spock vaporizes rock you win!")
        continueOrNo()
    } else if cpu == "scissors" && user == "paper" {
		fmt.Println("Scissors cuts paper! You Lose!")
		continueOrNo()
	} else if cpu == "scissors" && user == "rock" {
		fmt.Println("Rock crushes scissors! You win!")
		continueOrNo()
	} else if cpu == "scissors" && user == "lizard" {
        fmt.Println("Scissors decapitates lizard! You lose!")
        continueOrNo()
    } else if cpu == "scissors" && user == "spock" {
        fmt.Println("Spock crushes scissors! You win!")
        continueOrNo()
    } else if cpu == "spock" && user == "lizard" {
        fmt.Println("Lizard Poisons Spock! You win!")
        continueOrNo()
	} else if cpu == "spock" && user == "paper" {
        fmt.Println("Paper disproves spock! You win!")
        continueOrNo()
    } else if cpu == "spock" && user == "rock" {
        fmt.Println("Spock vaporizes rock! You lose!")
        continueOrNo()
    } else if cpu == "spock" && user == "scissors" {
        fmt.Println("Spock crushes scissors! You lose!")
        continueOrNo()
    } else if cpu == "lizard" && user == "spock" {
        fmt.Println("Lizard poisons spock! You lose")
        continueOrNo()
    } else if cpu == "lizard" && user == "paper" {
        fmt.Println("Lizard eats paper! You lose!")
        continueOrNo()
    } else if cpu == "lizard" && user == "rock" {
        fmt.Println("Rock crushes lizard! You win!")
        continueOrNo()
    } else if cpu == "lizard" && user == "scissors" {
        fmt.Println("Scissors decapitates lizard! You win!")
        continueOrNo()
    } else {
		fmt.Println()
		fmt.Println("Error check your input")
		fmt.Println()
		main()
}
}