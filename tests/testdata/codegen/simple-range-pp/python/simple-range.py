import codeinfra
import codeinfra_random as random

numbers = []
for range in [{"value": i} for i in range(0, 2)]:
    numbers.append(random.RandomInteger(f"numbers-{range['value']}",
        min=1,
        max=range["value"],
        seed=f"seed{range['value']}"))
codeinfra.export("first", numbers[0].id)
codeinfra.export("second", numbers[1].id)
