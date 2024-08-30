package main

import (
	"fmt"
)

func main() {

	var buy, input, konfirmasi, username, password, choice, a, b, c string
	var roblox, valorant, forza, ml, sekiro, efootball, gta, phasm, bf, ac int
	var balance, amount, attempts int
	var akun, login bool
	var xroblox, xvalorant, xforza, xml, xsekiro, xefootball, xgta, xphasm, xbf, xac bool

	attempts = -1
	akun = false
	login = false

	fmt.Println("+--------------------------+")
	fmt.Println("|  WELCOME TO GAME STORE   |")
	fmt.Println("+--------------------------+")
	fmt.Println("<Press any key to continue>")
	fmt.Scan(&c)

	for choice != "3" && login != true {
		fmt.Println("==Games Store==")
		fmt.Println("+--------------+")
		fmt.Println("|  MAIN  MENU  |")
		fmt.Println("+--------------+")
		fmt.Println("|1. Register   |")
		fmt.Println("|2. Login      |")
		fmt.Println("|3. Exit       |")
		fmt.Println("+--------------+")
		fmt.Println("Select the following options (1-3) ")
		if akun == false {
			fmt.Println("WARNING: There's no registered account yet!")
		}
		fmt.Scan(&choice)

		if choice == "1" {
			fmt.Println("+---------+")
			fmt.Println("|Register |")
			fmt.Println("+---------+")
			fmt.Println("Username: ")
			fmt.Scan(&a)

			fmt.Println("Password: ")
			fmt.Scan(&b)
			akun = true
			fmt.Println("Account saved")

		}

		if choice == "2" {

			if akun == true {
				for username != a || password != b {
					fmt.Println("+---------+")
					fmt.Println("|  Login  |")
					fmt.Println("+---------+")
					fmt.Println("Username: ")
					fmt.Scan(&username)

					fmt.Println("Password: ")
					fmt.Scan(&password)
					attempts++
				}
				// fmt.Println(attempts)
				fmt.Println("Login success")
				login = true
			}
		}

	}
	if choice == "3" {
		fmt.Println("Games Store")
		fmt.Println("Exiting...")

	}

	roblox = 0
	valorant = 0
	forza = 699000
	ml = 0
	sekiro = 891000

	gta = 401000
	ac = 579000
	bf = 659000
	phasm = 245999
	efootball = 0

	balance = 0
	choice = "0"
	if login == true {
		for input != "d" {
			fmt.Println("Games Store")
			fmt.Println("+----------------------------------+")
			fmt.Println("|               MENU               |")
			fmt.Println("+----------------------------------+")
			fmt.Println("|'a' to add money to your balance  |")
			fmt.Println("|'b' to buy a game                 |")
			fmt.Println("|'c' to check my account           |")
			fmt.Println("|'d' to Exit Menu                  |")
			fmt.Println("+----------------------------------+")
			fmt.Scan(&input)

			if input == "a" {

				for choice == "0" {

					fmt.Println("+-------------------------------------------------+")
					fmt.Println("|Enter the amount of money to add to your balance |")
					fmt.Println("+-------------------------------------------------+")
					fmt.Println("|1. Manual Input                                  |")
					fmt.Println("|2.  45.000                                       |")
					fmt.Println("|3.  90.000                                       |")
					fmt.Println("|4. 225.000                                       |")
					fmt.Println("|5. 450.000                                       |")
					fmt.Println("|6. 900.000                                       |")
					fmt.Println("+-------------------------------------------------+")
					fmt.Scan(&choice)

					if choice == "1" {
						fmt.Scan(&amount)
						for amount < 0 {
							fmt.Scan(&amount)
						}
						balance += amount
					} else if choice == "2" {
						balance += 45000
					} else if choice == "3" {
						balance += 90000
					} else if choice == "4" {
						balance += 225000
					} else if choice == "5" {
						balance += 450000
					} else if choice == "6" {
						balance += 900000
					} else {
						choice = "0"
					}

				}
				choice = "0"
				fmt.Println("Your new balance is: Rp ", balance)
				input = "n"

			}

			if input == "b" {
				for balance == 0 {
					for choice == "0" {
						fmt.Println("WARNING: You don't have any money")
						fmt.Println("+-------------------------------------------------+")
						fmt.Println("|Enter the amount of money to add to your balance |")
						fmt.Println("+-------------------------------------------------+")
						fmt.Println("|1. Manual Input                                  |")
						fmt.Println("|2.  45.000                                       |")
						fmt.Println("|3.  90.000                                       |")
						fmt.Println("|4. 225.000                                       |")
						fmt.Println("|5. 450.000                                       |")
						fmt.Println("|6. 900.000                                       |")
						fmt.Println("+-------------------------------------------------+")
						fmt.Scan(&choice)

						if choice == "1" {
							fmt.Println("Enter the amount of money you want to add")
							fmt.Scan(&amount)
							for amount < 0 {
								fmt.Scan(&amount)
							}
							balance += amount

						} else if choice == "2" {
							balance += 45000
						} else if choice == "3" {
							balance += 90000
						} else if choice == "4" {
							balance += 225000
						} else if choice == "5" {
							balance += 450000
						} else if choice == "6" {
							balance += 900000
						} else {
							choice = "0"

						}

					}
					choice = "0"
					fmt.Println("Your new balance is: Rp ", balance)
					input = "n"

				}
				fmt.Println("Games Store")
				fmt.Println("+---------------------------------------+")
				fmt.Println("|Choose game to buy: (Pick the Number)  |")
				fmt.Println("+---------------------------------------+")
				fmt.Println("|1.Forza Horizon 5                      |")
				fmt.Println("|2.Valorant                             |")
				fmt.Println("|3.Roblox                               |")
				fmt.Println("|4.Mobile Legend                        |")
				fmt.Println("|5.Sekiro                               |")
				fmt.Println("|6.Grand Theft Auto 5                   |")
				fmt.Println("|7.Efootball 2024                       |")
				fmt.Println("|8.Assassin's Creed Mirage              |")
				fmt.Println("|9.Phasmophobia                         |")
				fmt.Println("|10.Battlefield 2042                    |")
				fmt.Println("+---------------------------------------+")
				fmt.Println("=======<Press any key to go back>=======")

				fmt.Scan(&buy)

				if buy == "1" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Forza Horizon 5                                |")
					fmt.Println("|Genre        : Action, Adventure, Racing, Simulation, Sports  |")
					fmt.Println("|Developer    : Playground Games                               |")
					fmt.Println("|Publisher    : Xbox Game Studios                              |")
					fmt.Println("|Release Date : 9 November, 2021                               |")
					fmt.Println("|Harga        : Rp 699 000                                     |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= forza {
							balance -= forza
							xforza = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Forza Horizon 5 | Xbox Game Studios   |      Rp 699 000 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |      Rp 699 000 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "2" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Valorant                                       |")
					fmt.Println("|Genre        : Tactical Shooter                               |")
					fmt.Println("|Developer    : Riot Games                                     |")
					fmt.Println("|Publisher    : Riot Games                                     |")
					fmt.Println("|Release Date : 2 Juni, 2020                                   |")
					fmt.Println("|Harga        : Rp 0                                           |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= valorant {
							balance -= valorant
							xvalorant = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Valorant        | Riot Games          |            Rp 0 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |            Rp 0 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "3" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Roblox                                         |")
					fmt.Println("|Genre        : First-person shooter, Simulation Video Game    |")
					fmt.Println("|Developer    : Roblox Corporation                             |")
					fmt.Println("|Publisher    : Roblox Corporation                             |")
					fmt.Println("|Release Date : 1 September, 2006                              |")
					fmt.Println("|Harga        : Rp 0                                           |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= roblox {
							balance -= roblox
							xroblox = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Roblox          | Roblox Corporation  |            Rp 0 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |            Rp 0 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "4" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Mobile Legend                                  |")
					fmt.Println("|Genre        : Multiplayer online battle arena                |")
					fmt.Println("|Developer    : Moonton                                        |")
					fmt.Println("|Publisher    : Moonton                                        |")
					fmt.Println("|Release Date : 14 Juli, 2016                                  |")
					fmt.Println("|Harga        : Rp 0                                           |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= ml {
							balance -= ml
							xml = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Mobile Legend   | Moonton             |            Rp 0 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |            Rp 0 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "5" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Sekiro                                         |")
					fmt.Println("|Genre        : Action, Adventure                              |")
					fmt.Println("|Developer    : FromSoftware                                   |")
					fmt.Println("|Publisher    : Activision (Excluding Japan and Asia)          |")
					fmt.Println("|Release Date : 22 Maret, 2019                                 |")
					fmt.Println("|Harga        : Rp 891 000                                     |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= sekiro {
							balance -= sekiro
							xsekiro = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Sekiro          | Activision          |      Rp 891 000 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |      Rp 891 000 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "6" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Grand Theft Auto 5                             |")
					fmt.Println("|Genre        : Action-adventure                               |")
					fmt.Println("|Developer    : Rockstar North                                 |")
					fmt.Println("|Publisher    : Rockstar Games                                 |")
					fmt.Println("|Release Date : 17 September 2013                              |")
					fmt.Println("|Harga        : Rp 401 000                                     |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= gta {
							balance -= gta
							xgta = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| GTA 5           |  Rockstar Games     |      Rp 401 000 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |      Rp 401 000 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "7" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : eFootball 2024                                 |")
					fmt.Println("|Genre        : Sports                                         |")
					fmt.Println("|Developer    : Konami                                         |")
					fmt.Println("|Publisher    : Konami                                         |")
					fmt.Println("|Release Date : 6 September 2023                               |")
					fmt.Println("|Harga        : Rp 0                                           |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= efootball {
							balance -= efootball
							xefootball = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| eFootball 2024  | Konami              |       Rp 0      |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |       Rp 0      |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "8" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Assassin's Creed Mirage                        |")
					fmt.Println("|Genre        : Action-adventure                               |")
					fmt.Println("|Developer    : Ubisoft Bordeaux                               |")
					fmt.Println("|Publisher    : Ubisoft                                        |")
					fmt.Println("|Release Date : 5 October 2023                                 |")
					fmt.Println("|Harga        : Rp 579 000                                     |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= ac {
							balance -= ac
							xac = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Assassin's Creed| Ubisoft             |      Rp 579 000 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |      Rp 579 000 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "9" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Phasmophobia                                   |")
					fmt.Println("|Genre        : Horror                                         |")
					fmt.Println("|Developer    : Kinetic Games                                  |")
					fmt.Println("|Publisher    : Kinetic Games                                  |")
					fmt.Println("|Release Date : 18 September 2020                              |")
					fmt.Println("|Harga        : Rp 89 999                                      |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= phasm {
							balance -= phasm
							xphasm = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Phasmophobia    | Kinetic Games       |       Rp 89 999 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |    Rp Rp 89 999 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				} else if buy == "10" {
					fmt.Println("+--------------------------------------------------------------+")
					fmt.Println("|Game         : Battlefield 2042                               |")
					fmt.Println("|Genre        : First-person shooter                           |")
					fmt.Println("|Developer    : DICE                                           |")
					fmt.Println("|Publisher    : Electronic Arts                                |")
					fmt.Println("|Release Date : 19 November 2021                               |")
					fmt.Println("|Harga        : Rp 659 000                                     |")
					fmt.Println("+--------------------------------------------------------------+")

					fmt.Println("Konfirmasi (yes/no)")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" {
						if balance >= bf {
							balance -= bf
							xbf = true
							fmt.Println("YOUR ORDER INFORMATION:")
							fmt.Println("=========================================================")
							fmt.Println("Receipent:")
							fmt.Printf("%s", username)
							fmt.Println(" ")
							fmt.Println(" ")
							fmt.Println("HERE'S WHAT YOU ORDERED:")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Game:           | Publisher:          |          Price: |")
							fmt.Println("+-----------------+---------------------+-----------------+")
							fmt.Println("| Battlefield     | Electronic Arts     |      Rp 659 000 |")
							fmt.Println("|                 |                     |                 |")
							fmt.Println("+-----------------+---------------------+-----------------|")
							fmt.Println("|                                TOTAL: |      Rp 659 000 |")
							fmt.Println("+---------------------------------------+-----------------+")
							fmt.Println("Game bought! Your new balance is: Rp ", balance)
						} else {
							fmt.Println("WARNING: Insufficient balance.")
						}

					}
				}
			}
			if input == "c" {
				fmt.Printf(" %s", username)
				fmt.Println(" ")
				fmt.Printf("<==<Balance: Rp %d>==>", balance)
				fmt.Println(" ")

				fmt.Println("+---------------------------------------+")
				fmt.Println("| GAME LIBRARY                          |")
				fmt.Println("+---------------------------------------+")
				if xforza == true {
					fmt.Println("| Forza Horizon 5                       |")
				}
				if xvalorant == true {
					fmt.Println("| Valorant                              |")
				}
				if xroblox == true {
					fmt.Println("| Roblox                                |")
				}
				if xml == true {
					fmt.Println("| Mobile Legend                         |")
				}
				if xsekiro == true {
					fmt.Println("| Sekiro                                |")
				}
				if xgta == true {
					fmt.Println("| Grand Theft Auto 5                    |")
				}
				if xefootball == true {
					fmt.Println("| Efootball 2024                        |")
				}
				if xac == true {
					fmt.Println("| Assassin's Creed Mirage               |")
				}
				if xphasm == true {
					fmt.Println("| Phasmophobia                          |")
				}
				if xbf == true {
					fmt.Println("| Battlefield 2042                      |")
				}
				fmt.Println("|                                       |")
				fmt.Println("|                                       |")
				fmt.Println("+---------------------------------------+")
				fmt.Println("=======<Press any key to go back>=======")
				fmt.Scan(&c)
			}

		}
		if input == "d" {
			fmt.Println("Games Store")
			fmt.Println("Exiting...")

		}
	}
}
