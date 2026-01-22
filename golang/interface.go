package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	eat, move, speak string
}

func (c *Cow) Eat()   { fmt.Println(c.eat) }
func (c *Cow) Move()  { fmt.Println(c.move) }
func (c *Cow) Speak() { fmt.Println(c.speak) }

type Bird struct {
	eat, move, speak string
}

func (b *Bird) Eat()   { fmt.Println(b.eat) }
func (b *Bird) Move()  { fmt.Println(b.move) }
func (b *Bird) Speak() { fmt.Println(b.speak) }

type Snake struct {
	eat, move, speak string
}

func (s *Snake) Eat()   { fmt.Println(s.eat) }
func (s *Snake) Move()  { fmt.Println(s.move) }
func (s *Snake) Speak() { fmt.Println(s.speak) }

var animals = make(map[string]Animal)

func newAnimal(name, kind string) error {
	switch kind {
	case "cow":
		animals[name] = &Cow{"grass", "walk", "moo"}
	case "bird":
		animals[name] = &Bird{"worms", "fly", "peep"}
	case "snake":
		animals[name] = &Snake{"mice", "slither", "hsss"}
	default:
		return fmt.Errorf("unknown animal type")
	}
	fmt.Println("Created It")
	return nil
}

func query(name, action string) error {
	animal, ok := animals[name]
	if !ok {
		return fmt.Errorf("animal not found")
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		return fmt.Errorf("unknown action")
	}
	return nil
}

func userOption(option string) (func(string, string) error, error) {
	switch option {
	case "newanimal":
		return newAnimal, nil
	case "query":
		return query, nil
	default:
		return nil, fmt.Errorf("invalid command")
	}
}

func main() {
	var cmd, arg1, arg2 string

	for {
		fmt.Print("> ")
		fmt.Scan(&cmd, &arg1, &arg2)

		fn, err := userOption(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if err := fn(arg1, arg2); err != nil {
			fmt.Println(err)
		}
	}
}
