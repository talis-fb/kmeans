import numpy as np

# Number of clusters
k = 5
# Number of points per cluster
n = 100

# Generate random cluster centers with integer coordinates
centers = np.random.randint(0, 1000, (k, 2))

# Generate random points around the cluster centers with integer coordinates
points = []
for center in centers:
    points.extend(np.random.randint(center - 100, center + 100, (n, 2)))

# Convert to NumPy array
points = np.array(points, dtype=int)

# Shuffle the points
np.random.shuffle(points)

# Save the points to a CSV file
np.savetxt('large_dataset.csv', points, delimiter=' ', fmt='%d')
