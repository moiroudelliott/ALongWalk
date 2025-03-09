package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// Point représente les coordonnées d'une cellule dans la grille.
type Point struct {
	x, y int
}

// Edge représente une connexion entre deux nœuds (jonctions) avec un poids (distance).
type Edge struct {
	to, weight int
}

var (
	globalBest int      // Meilleure longueur de chemin trouvée jusqu'à présent.
	totalNodes int      // Nombre total de nœuds (jonctions) dans le graphe.
	maxEdge    int      // Poids maximal d'une arête dans le graphe.
	adj        [][]Edge // Liste d'adjacence indexée pour le graphe des jonctions.
)

// readGridFromFile charge la grille à partir d'un fichier.
func readGridFromFile(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Erreur d'ouverture du fichier:", err)
		os.Exit(1)
	}
	defer file.Close()

	var grid [][]byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	return grid
}

// isJunction détermine si une cellule (x,y) est une jonction.
func isJunction(grid [][]byte, x, y, rows, cols int) bool {
	if grid[x][y] == '#' {
		return false
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	count := 0
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] != '#' {
			count++
		}
	}
	return count > 2 || x == 0 || x == rows-1
}

// computeJunctions précalcule quelles cellules sont des jonctions.
func computeJunctions(grid [][]byte) [][]bool {
	rows, cols := len(grid), len(grid[0])
	junctions := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		junctions[i] = make([]bool, cols)
		for j := 0; j < cols; j++ {
			junctions[i][j] = isJunction(grid, i, j, rows, cols)
		}
	}
	return junctions
}

// index convertit les coordonnées 2D en un index linéaire.
func index(x, y, cols int) int {
	return x*cols + y
}

// explorePath explore récursivement la grille à partir d'une jonction jusqu'à atteindre une autre jonction.
func explorePath(grid [][]byte, junctions [][]bool, start Point, x, y, steps int, visited []bool, rows, cols int, graph map[Point]map[Point]int) {
	idx := index(x, y, cols)
	if visited[idx] {
		return
	}
	visited[idx] = true

	if junctions[x][y] && (x != start.x || y != start.y) {
		end := Point{x, y}
		graph[start][end] = steps
		graph[end][start] = steps
		visited[idx] = false
		return
	}

	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, d := range dirs {
		nx, ny := x+d[0], y+d[1]
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] != '#' {
			explorePath(grid, junctions, start, nx, ny, steps+1, visited, rows, cols, graph)
		}
	}
	visited[idx] = false
}

// buildGraph construit un graphe où les nœuds sont des jonctions et les arêtes représentent les chemins entre elles.
func buildGraph(grid [][]byte) map[Point]map[Point]int {
	rows, cols := len(grid), len(grid[0])
	junctions := computeJunctions(grid)

	graph := make(map[Point]map[Point]int)
	var junctionPoints []Point

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if junctions[x][y] {
				p := Point{x, y}
				junctionPoints = append(junctionPoints, p)
				graph[p] = make(map[Point]int)
			}
		}
	}

	visited := make([]bool, rows*cols)
	for _, start := range junctionPoints {
		explorePath(grid, junctions, start, start.x, start.y, 0, visited, rows, cols, graph)
	}

	return graph
}

// buildIndexedGraph convertit le graphe des jonctions en un graphe indexé.
func buildIndexedGraph(graph map[Point]map[Point]int) ([]Point, map[Point]int, [][]Edge) {
	nodes := make([]Point, 0, len(graph))
	pointToIndex := make(map[Point]int)
	for p := range graph {
		pointToIndex[p] = len(nodes)
		nodes = append(nodes, p)
	}

	adjList := make([][]Edge, len(nodes))
	for p, edges := range graph {
		u := pointToIndex[p]
		for neighbor, weight := range edges {
			v := pointToIndex[neighbor]
			adjList[u] = append(adjList[u], Edge{to: v, weight: weight})
		}
		sort.Slice(adjList[u], func(i, j int) bool {
			return adjList[u][i].weight > adjList[u][j].weight
		})
	}

	return nodes, pointToIndex, adjList
}

// dfs explore le graphe à partir du nœud u avec un masque de visite 'mask', une somme courante 'curr', et 'count' nœuds visités.
func dfs(u int, mask uint64, curr, count int) int {
	rem := totalNodes - count
	if curr+rem*maxEdge <= globalBest {
		return curr
	}

	bestLocal := curr
	for _, edge := range adj[u] {
		if mask&(1<<edge.to) == 0 {
			candidate := dfs(edge.to, mask|(1<<edge.to), curr+edge.weight, count+1)
			if candidate > bestLocal {
				bestLocal = candidate
			}
		}
	}

	if bestLocal > globalBest {
		globalBest = bestLocal
	}
	return bestLocal
}

func algo(filename ...string) int {

	// Demander à l'utilisateur le chemin du fichier.
	var filePath string

	// Vérification du paramètre optionnel
	if len(filename) > 0 && filename[0] != "" {
		filePath = filename[0]
	} else {
		// Demande à l'utilisateur le chemin du fichier si non fourni
		fmt.Print("Veuillez entrer le chemin du fichier d'entrée: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		filePath = scanner.Text()
	}

	// Lecture de la grille.
	grid := readGridFromFile(filePath)
	rows, cols := len(grid), len(grid[0])

	// Construction du graphe des jonctions.
	graph := buildGraph(grid)

	// Conversion en graphe indexé.
	nodes, pointToIndex, indexedAdj := buildIndexedGraph(graph)
	adj = indexedAdj
	totalNodes = len(nodes)

	// Calcul du poids maximal d'une arête.
	maxEdge = 0
	for _, edges := range adj {
		for _, e := range edges {
			if e.weight > maxEdge {
				maxEdge = e.weight
			}
		}
	}

	// Détermination de l'entrée et de la sortie.
	var start, end Point
	for y := 0; y < cols; y++ {
		if grid[0][y] == '.' {
			start = Point{0, y}
		}
		if grid[rows-1][y] == '.' {
			end = Point{rows - 1, y}
		}
	}

	// Vérifie si l'entrée et la sortie sont des jonctions.
	if _, ok := pointToIndex[start]; !ok {
		fmt.Println("Erreur : l'entrée n'est pas une jonction.")
		return -1
	}
	if _, ok := pointToIndex[end]; !ok {
		fmt.Println("Erreur : la sortie n'est pas une jonction.")
		return -1
	}

	// Initialisation de la meilleure solution globale.
	globalBest = 0

	// Exécution de DFS en parallèle à partir de l'entrée.
	maxTotal := dfs(pointToIndex[start], 1<<pointToIndex[start], 0, 1)

	return maxTotal
}
