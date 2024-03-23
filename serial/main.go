package main

import (
	"fmt"
	// "math"
)

import . "github.com/talis-fb/kmeans/common"


func main() {
	data := []Point{
		{1, 2},
		{5, 8},
		{1, 3},
		{6, 9},
		{2, 3},
		{5, 10},
		{1, 1},
		{7, 8},
	}

	k := 2
	clusters := KMeans(data, k)
	for i, cluster := range clusters {
		fmt.Printf("Cluster %d:\n", i+1)
		fmt.Printf("Center: %v\n", cluster.Center)
		fmt.Printf("Points: %v\n", cluster.Points)
	}
}

// KMeans performs K Means clustering on the given data points.
func KMeans(data []Point, k int) []Cluster {
	clusters := initializeClusters(data, k)

	for {
		assignPoints(data, clusters)

		newCenters := calculateNewCenters(clusters)

		if converged(clusters, newCenters) {
			break
		}

		for i := range clusters {
			clusters[i].Center = newCenters[i]
		}
	}

	return clusters
}

// TODO: make this random
func initializeClusters(data []Point, k int) []Cluster {
  clusters := make([]Cluster, k)
	for i := range clusters {
		clusters[i].Center = data[i]
	}
	return clusters
  // return Map(data, func(point Point) Cluster {
  //   return Cluster{Center: point}
  // })
}

func assignPoints(data []Point, clusters []Cluster) {
	for i := range clusters {
		quantPointsInCluster := len(clusters[i].Points)
		if quantPointsInCluster > 0 {
			clusters[i].Points = make([]Point, 0)
		}
	}

	for _, point := range data {
		var nearestCluster *Cluster = getNearestCluster(point, clusters)

		nearestCluster.Points = append(nearestCluster.Points, point)
	}
}

func getNearestCluster(point Point, clusters []Cluster) *Cluster {
  clustersCenterPoints := Map(clusters, func(cluster Cluster) Point {
    return cluster.Center
  });

  distOfEachClusterCenterPoints := Map(clustersCenterPoints, func(clusterCenter Point) float64 {
    return point.EuclideanDistance(clusterCenter)
  })

  targetClusterIndex := IndexOfMin(distOfEachClusterCenterPoints)


	return &clusters[targetClusterIndex]

	// targetClusterIndex := 0

	// // Start with higher (infinity) and it getting lower
	// lowerDistanceCluster := math.Inf(1)
	//
	// targetClusterIndex := 0
	//
	// for i, cluster := range clusters {
	// 	dist := point.EuclideanDistance(cluster.Center)
	// 	if dist < lowerDistanceCluster {
	// 		lowerDistanceCluster = dist
	// 		targetClusterIndex = i
	// 	}
	// }
	//
	// return &clusters[targetClusterIndex]
}

func calculateNewCenters(clusters []Cluster) []Point {
  return Map(clusters, func(cluster Cluster) Point {
    return calculateCenter(&cluster)
  })
}

func calculateCenter(cluster *Cluster) Point {
	if len(cluster.Points) == 0 {
		return cluster.Center
	}

	sumX := 0.0
	sumY := 0.0
	for _, point := range cluster.Points {
		sumX += point.X
		sumY += point.Y
	}

	quantPointsInCluster := float64(len(cluster.Points))

	meanX := sumX / quantPointsInCluster
	meanY := sumY / quantPointsInCluster

	return Point{meanX, meanY}
}

func converged(clusters []Cluster, newCenters []Point) bool {
  for i, cluster := range clusters {
		if !equalPoints(cluster.Center, newCenters[i]) {
			return false
		}
	}
	return true

  // clustersCenters := Map(clusters, func(cluster Cluster) Point {
  //   return cluster.Center
  // })

  // a := All(clustersCenters, func(cluster []Cluster) bool {
  //   return equalPointList(cluster, newCenters)
  // })

  // return equalPointList(clustersCenters, newCenters)
}

func equalPointList(a, b []Point) bool {
  for i := range len(a) {
		if !equalPoints(a[i], b[i]) {
			return false
		}
	}
	return true
}

func equalPoints(a, b Point) bool {
	return a.X == b.X && a.Y == b.Y
}
