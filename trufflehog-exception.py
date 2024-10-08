import re
import json

# Load patterns from the JSON file
with open("truffleHogAllowRules.json", "r") as f:
    patterns_list = json.load(f)

# Compile the patterns into regex objects
patterns = [re.compile(pattern) for pattern in patterns_list]

# Function to determine if a block should be excluded
def should_exclude(block):
    for pattern in patterns:
        if any(pattern.search(line) for line in block):
            return True
    return False

# Read the input file
with open("trufflehog_output.text", "r") as f:
    lines = f.readlines()

# Process the file and remove matching blocks
output_lines = []
current_block = []

for line in lines:
    if line.startswith("Found unverified result"):
        if current_block and not should_exclude(current_block):
            output_lines.extend(current_block)
        current_block = [line]
    else:
        current_block.append(line)

# Append the last block if it doesn't match the patterns
if current_block and not should_exclude(current_block):
    output_lines.extend(current_block)

# Write the filtered output to a new file
with open("trufflehog_filtered_output.text", "w") as f:
    f.writelines(output_lines)

print("Filtered output saved to trufflehog_filtered_output.text")
