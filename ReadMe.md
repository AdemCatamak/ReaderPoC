# Key-Value Store Reader Implementation

## Overview
This project implements a key-value store with two distinct strategies for handling case-insensitive key comparison in the `Get` method:
1. **Comparator-based Strategy**: Uses a comparator interface to define custom comparison logic.
2. **Inheritance-based Strategy**: Creates separate reader implementations, each with its own `Get` method for specific comparison logic.

## Solution Structure

### 1. Comparator-based Strategy
Located in `comparisonStrategySolution`, this approach defines a comparator interface for key comparison, enabling flexible comparison methods by passing different comparator implementations to the reader. Key components:
- **`comparator.go`**: Defines the comparator interface and custom comparator functions.
- **`defaultReader.go`**: Implements a reader that accepts a comparator, allowing it to use various comparison strategies.

### 2. Inheritance-based Strategy
Found in `inheritanceSolution`, this solution involves creating multiple reader structs for different comparison types:
- **`defaultReader.go`**: A reader with case-sensitive `Get` behavior.
- **`caseInsensitiveReader.go`**: Extends `defaultReader` to support case-insensitive key matching.

## Usage
To use either approach, instantiate the appropriate reader or comparator, then call the `Get` method with the desired key.

1. **Comparator-based Reader**:
    - Create a `defaultReader` with a specific comparator.
    - Use the `Get` method to retrieve values based on the provided comparison logic.

2. **Inheritance-based Reader**:
    - Choose between `defaultReader` or `caseInsensitiveReader` for the desired key sensitivity.
    - Call `Get` to retrieve values, with comparison behavior defined by the reader type.
