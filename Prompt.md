# Prompt: Composite Design Pattern in Go

Implement the Composite Design Pattern in Go based on the lecture example.

Requirements:

1. Define a common interface `Item` with the methods `Price()` and `Display()`.
2. Implement the components `Window`, `Glass`, and `Frame`.
3. Allow a `Window` to contain multiple components.
4. Both individual components and composite objects must implement the same `Item` interface.
5. Implement `Price()` to calculate the total price recursively.
6. Implement `Display()` to display the complete component hierarchy.
7. Use idiomatic and well-structured Go code.
8. Include a `main.go` demonstrating the Composite Pattern.
9. Add unit tests for the price calculation and component hierarchy.
10. Add a README.md explaining the architecture and how to run the program and tests.
Keep the implementation simple and consistent with the lecture diagram. Do not introduce unnecessary abstractions or dependencies.