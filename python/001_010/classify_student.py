def calculate_avg_score(scores):
    avg = sum(scores) / len(scores)
    return round(avg, 2)


def classify_student(score):
    if score >= 8.5:
        return "Excellent"
    if score >= 7.0:
        return "Very good"
    if score >= 5.5:
        return "Good"
    if score >= 4.0:
        return "Average"
    return "Below average"


tests = [
    [10.0, 9.0, 4.8, 9, 8.0],
    [6.0, 7.0, 4.8, 9, 8.0],
    [8.5, 5.8, 8.8, 9.2, 6.0],
    [6.6, 9.0, 6.8, 9.2, 7.0],
]

for scores in tests:
    avg_score = calculate_avg_score(scores)
    classification = classify_student(avg_score)
    print("Scores:", scores, "average:", avg_score, "classification:", classification)
