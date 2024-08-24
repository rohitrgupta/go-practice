package main

func Hello(name string, language string) string {
  if name == "" {
    name = "World"
  }
  if language == "Spanish" {
    return "Hola, " + name
  }
  if language == "French" {
    return "Bonjour, " + name
  }
  return "Hello, " + name
  
}
