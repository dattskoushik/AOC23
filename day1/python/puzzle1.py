def sum_first_last_digits(file_path):
    # Initialize sum to zero
    total_sum = 0
    # Open the file and read it line by line
    with open(file_path, 'r') as file:
        for line in file:
            # Initialize variables to store the first and last numbers
            first_digit = None
            last_digit = None
            final_num = None

            # Traverse the line to find the first and last numbers
            for char in line:
                if char.isdigit() and first_digit is None:
                    first_digit = char
                    break
            for char in reversed(line):
                if char.isdigit() and last_digit is None:
                    last_digit = char
                    break
            final_num = first_digit + last_digit            
            # Add the first and last numbers to the total sum
            if final_num is not None:
                total_sum += int(final_num)

    # Return the final sum
    return total_sum

def main():
    # Get the file path from the user
    file_path = input("Enter the path to your text file: ")

    # Calculate and print the sum
    result = sum_first_last_digits(file_path)
    print("Sum of first and last numbers:", result)

if __name__ == "__main__":
    main()
