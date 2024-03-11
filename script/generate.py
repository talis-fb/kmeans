import numpy as np

# Set the number of clusters
num_clusters = 5

# Set the number of points per cluster
points_per_cluster = 100_000

# Set the dimensions of each point
num_dimensions = 2

# Generate random cluster centers
cluster_centers = np.random.rand(num_clusters, num_dimensions) * 100

# Generate random points around the cluster centers
data = np.concatenate([np.random.randn(points_per_cluster, num_dimensions) + center for center in cluster_centers])

# Optionally, save the data to a file
np.savetxt('huge_dataset.csv', data, delimiter=',')
