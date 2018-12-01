# AdventOfCode2018
Another year has passed, I wanted to try out yet another language, so I'm using **go** this year. This is actually my first time writing go, so hopefully it will improve as days pass.

The day before I did some preparation to set up the repository, basic functions that I used a lot previous years and some basic 
stuff to get familiarised with the language. You can check those out in `common.go`

##Day 1

Let's go! After the standard introduction, we get to the point. We will be given a list of integers
with their sign. We start from 0 and need to add all the given numbers to obtain the final result. _I'm already missing python..._

Before thinking about the solution, let's get on with the tests:

```go
func TestFrequencyCalculator(t *testing.T) {
	tables := []struct {
		input []string
		output int
	}{
		{[]string{"+1", "-2", "+3", "+1"}, 3},
		{[]string{"+1", "+1", "+1"}, 3},
		{[]string{"+1", "+1", "-2"}, 0},
		{[]string{"-1", "-2", "-3"}, -6},
	}

	for _, table := range tables {
		result, err := FrequencyCalculator(table.input)
		if err != nil {
			log.Fatal(err)
			return
		}
		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}
```

As usual, we get some nice failures, let's fix them!  
The solution for the problem is just an accumulator, but to be able to perform operations on the numbers, first we need to 
turn them into signed integers. I made a function that takes the frequencies and outputs a list of `int`:
```go
func readFrequencies(input []string) ([]int, error) {
	var symbol = map[string]int{
		"+": 1,
		"-": -1,
	}

	var frequencies []int
	for _, element := range input {
		multiplier := symbol[string(element[0])]
		value, err := strconv.Atoi(element[1:])
		if err != nil {
			return nil, err
		}

		frequencies = append(frequencies, value*multiplier)
	}
	return frequencies, nil
}
```

I used a dictionary to map the sign symbol with their numeric value. As I'm writing this I'm thinking that it seems like a nice
bit of code to refactor, I'll extract into the common library if I have to parse some signed integers again.

The code for the accumulator is pretty straight forward:
```go
func FrequencyCalculator(input []string) (int, error) {
	frequencies, err := readFrequencies(input)
	if err != nil {
		return 0, err
	}

	currentFrequency := 0
	for _, element := range frequencies {
		currentFrequency += element
	}
	return currentFrequency, nil
}
```

All the test pass now! We can now run today's runner:

```go
func runDay01(){
	fmt.Println("Read input 1!")
	input, err := readDailyInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := day01.FrequencyCalculator(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for part 1 is: %d", result) // returns 582
}
```

That's it for the first star, now the second involves treating the list as a circular list.  
These are the unit tests:
```go
func TestRepeatedFrequenciesCalculator(t *testing.T) {
	tables := []struct {
		input []string
		output int
	}{
		{[]string{"+1", "-2", "+3", "+1"}, 2},
		{[]string{"+1", "-1"}, 0},
		{[]string{"+3", "+3", "+4", "-2", "-4"}, 10},
		{[]string{"-6", "+3", "+8", "+5", "-6"}, 5},
		{[]string{"+7", "+7", "-2", "-7", "-4"}, 14},
	}

	for _, table := range tables {
		result, err := RepeatedFrequenciesCalculator(table.input)
		if err != nil {
			log.Fatal(err)
			return
		}
		if result != table.output {
			t.Errorf("Got %d, expected %d", result, table.output)
		}
	}
}
```
I would love to reuse some of the code from the first implementation, but it seems that it won't be possible...  
I had heard about `Ring` data type in _go_, to I went directly for it. It might not be the simplest solution, 
but it will do the job. It seems like rings are cursor-like structures where you ask for the next position using `.Next()`.
Here is my implementation of the solution:
```go
func RepeatedFrequenciesCalculator(input []string) (int, error) {
	frequencies, err := readFrequencies(input)
	if err != nil {
		return 0, err
	}
	ringFreq := ring.New(len(frequencies))
	for i := 0; i < ringFreq.Len(); i++ {
		ringFreq.Value = frequencies[i]
		ringFreq = ringFreq.Next()
	}

	currentFrequency := 0
	var freqStore = map[int]int{
		0: 1,
	}
	for {
		currentFrequency += ringFreq.Value.(int)
		if freqStore[currentFrequency] >= 1 {
			return currentFrequency, nil
		}
		freqStore[currentFrequency] += 1
		ringFreq = ringFreq.Next()
	}
}
```

The first part after converting the frequencies to signed integers is just initialising the ring. The main
logic uses an infinite for loop that stores any seen frequency on a dictionary. If the frequency has been seen before it
will return that frequency as the first one to be seen twice.

After checking the tests, all of them pass, let's now modify the main runner to see the result:
```go
func runDay01(){
	fmt.Println("Read input 1!")
	input, err := readDailyInput(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := day01.FrequencyCalculator(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for part 1 is: %d", result) // returns 582

	fmt.Println()

	result, err = day01.RepeatedFrequenciesCalculator(input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Result for part 2 is: %d", result) // returns 488
}
```

Day 1 is over! Easy problem as usual, but it took me a decent amount of time since I had to learn __go__
on the __go__ (no pun intended). I might revisit it as I learn a few more tricks!
