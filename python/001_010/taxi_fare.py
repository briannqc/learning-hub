def calculate_taxi_fare(distance):
    if distance <= 1.0:
        return 10
    if distance <= 10.0:
        return 10 + 8.5 * (distance - 1)
    return 10 + 8.5 * 9 + 7.5 * (distance - 10)


tests = [0.5, 1, 2, 5, 9.9, 10, 11, 15, 20]
for t in tests:
    print("Fare for", t, "km is", calculate_taxi_fare(t))
