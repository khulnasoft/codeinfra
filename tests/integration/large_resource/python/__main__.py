import codeinfra

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Create a very deep array (>100 levels)
deep_array = []
for i in range(0, 200):
    deep_array = [deep_array]

codeinfra.export("long_string",  long_string)
codeinfra.export("deep_array",  deep_array)
