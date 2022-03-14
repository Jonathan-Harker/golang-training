# Golang Training

<details>
<summary>Exercise 1 - Pass The Int</summary>

Write a program where two goroutines pass an integer back and forth ten times.  
Display when each goroutine receives the integer.  
Increment the integer with each pass.  
Once the integer equals ten, terminate the program cleanly.  
[Example Here](/concurancy/exercises/pass_the_int.go)
</details>

<details>
<summary>Exercise 2 - Generate 100 random numbers concurrently</summary>

Write a program that uses a fan out pattern to generate 100 random numbers concurrently.   
Have each goroutine generate a single random number and return that number to the main goroutine over a buffered channel.   
Set the size of the buffer channel so no send every blocks.   
Don't allocate more buffers than you need.   
Have the main goroutine display each random number is receives and then terminate the program.
[Example Here](/concurancy/exercises/generate_100_numbers.go)
</details>

<details>
<summary>Exercise 3 - Generate random odd numbers concurrently</summary>

Write a program that uses goroutines to generate up to 100 random numbers.  
Do not send values that are divisible by 2.  
Have the main goroutine receive values and add them to a slice.
[Example Here](/concurancy/exercises/generate_random_odd_nums.go)
</details>

<details>
<summary>Exercise 4 - Generate 100 random odd numbers concurrently</summary>

Write a program that creates a fixed set of workers to generate random numbers.   
Discard any number divisible by 2.   
Continue receiving until 100 numbers are received.   
Tell the workers to shut down before terminating.
[Example Here](/concurancy/exercises/generate_100_random_odd_nums.go)
</details>

<details>
<summary>Bench Testing</summary>

`cd testing/testing_exercises`  
`go test -run none -bench . -benchtime 3s -benchmem`

Write three benchmark tests for converting an integer into a string.  
Test the following functions: 
* fmt.Sprintf
* strconv.FormatInt 
* strconv.Itoa

Identify which function performs the best.

Results:

| Name                     | Amount Run | Time per op | Mem per op | Allocations |
|--------------------------|------------|-------------|------------|-------------|
| BenchmarkSprintfTest-8   | 53442950   | 64.45 ns/op | 3 B/op     | 1 allocs/op |
| BenchmarkFormatIntTest-8 | 183734679  | 19.75 ns/op | 3 B/op     | 1 allocs/op |
| BenchmarkItoaTest-8      | 181195120  | 19.70 ns/op | 3 B/op     | 1 allocs/op |

</details>

<details>
<summary>SHA 256</summary>

Write a program that is given a list of file names as arguments.  
This then prints the sha256 sum for the contents of each file.  
Print the hashes as a hex string.
</details>
