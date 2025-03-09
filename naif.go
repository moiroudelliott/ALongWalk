package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type PointNaive struct {
	x, y int
}

// Lire la grille depuis un fichier
func readGridFromFileNaive(filename string) [][]byte {
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

// Vérifier si une cellule est un carrefour (sommet)
func isJunctionNaive(grid [][]byte, x, y int) bool {
	if grid[x][y] == '#' {
		return false
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	count := 0

	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] != '#' {
			count++
		}
	}

	// Un carrefour a au moins 3 directions valides ou est sur le bord (entrée/sortie)
	return count > 2 || x == 0 || x == len(grid)-1
}

// Construire le graphe
func buildGraphNaive(grid [][]byte) map[PointNaive]map[PointNaive]int {
	graph := make(map[PointNaive]map[PointNaive]int)
	junctions := []PointNaive{}

	// Détecter tous les carrefours
	for x := range grid {
		for y := range grid[0] {
			if isJunctionNaive(grid, x, y) {
				junctions = append(junctions, PointNaive{x, y})
				graph[PointNaive{x, y}] = make(map[PointNaive]int)
			}
		}
	}

	// Explorer les chemins et relier les carrefours
	for _, start := range junctions {
		visited := make(map[PointNaive]bool)
		explorePathNaive(grid, start, start.x, start.y, 0, visited, graph)
	}

	return graph
}

// Explorer un chemin récursivement jusqu'à un carrefour
func explorePathNaive(grid [][]byte, start PointNaive, x, y, steps int, visited map[PointNaive]bool, graph map[PointNaive]map[PointNaive]int) {
	if visited[PointNaive{x, y}] {
		return
	}
	visited[PointNaive{x, y}] = true

	if isJunctionNaive(grid, x, y) && (x != start.x || y != start.y) {
		end := PointNaive{x, y}
		graph[start][end] = steps
		graph[end][start] = steps
		return
	}

	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] != '#' {
			explorePathNaive(grid, start, nx, ny, steps+1, visited, graph)
		}
	}

	visited[PointNaive{x, y}] = false
}

// **DFS stable pour trouver le plus long chemin**
func dfsLongestPath(graph map[PointNaive]map[PointNaive]int, node PointNaive, visited map[PointNaive]bool, steps int, maxSteps *int) {
	visited[node] = true

	// Met à jour la distance maximale atteinte
	if steps > *maxSteps {
		*maxSteps = steps
	}

	// **Trier les voisins pour assurer un ordre constant**
	neighbors := make([]PointNaive, 0, len(graph[node]))
	for neighbor := range graph[node] {
		neighbors = append(neighbors, neighbor)
	}
	sort.Slice(neighbors, func(i, j int) bool {
		return neighbors[i].x < neighbors[j].x || (neighbors[i].x == neighbors[j].x && neighbors[i].y < neighbors[j].y)
	})

	// Explorer les voisins triés
	for _, neighbor := range neighbors {
		if !visited[neighbor] {
			dfsLongestPath(graph, neighbor, visited, steps+graph[node][neighbor], maxSteps)
		}
	}

	// **Backtracking**
	visited[node] = false
}

func naive(filename ...string) {
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

	grid := readGridFromFileNaive(filePath) // Lecture de la grille

	graph := buildGraphNaive(grid)

	startingPointNaives := []PointNaive{}
	for y := 0; y < len(grid[0]); y++ {
		if grid[0][y] == '.' {
			startingPointNaives = append(startingPointNaives, PointNaive{0, y})
		}
	}

	maxSteps := 0
	for _, start := range startingPointNaives {
		visited := make(map[PointNaive]bool)
		dfsLongestPath(graph, start, visited, 0, &maxSteps)
	}
}
