Voici votre README mis à jour, sans les exemples et la conclusion :  

---

# **Installation de Go et Exécution du Programme**

## **Installation de Go**

Pour exécuter ce programme, vous devez d'abord installer **Go (Golang)** sur votre machine. Voici les étapes pour l'installer :

1. **Téléchargez Go** :
   - Rendez-vous sur le site officiel de Go : [https://golang.org/dl/](https://golang.org/dl/).
   - Téléchargez la version correspondant à votre système d'exploitation (Windows, macOS, Linux).

2. **Installez Go** :
   - **Sur Windows** : Exécutez le fichier d’installation téléchargé et suivez les instructions.
   - **Sur macOS** : Ouvrez le fichier `.pkg` téléchargé et suivez les instructions.
   - **Sur Linux** : Extrayez l'archive téléchargée dans `/usr/local`, puis ajoutez `go/bin` à votre variable `PATH`.

3. **Vérifiez l'installation** :
   - Ouvrez un terminal et tapez la commande :
     ```bash
     go version
     ```
   - Si l’installation est correcte, la version de Go installée s'affichera.

---

## **Exécution du Programme**

### **1. Cloner le projet**
Si vous utilisez Git, clonez le dépôt contenant le code source :
```bash
git clone https://github.com/moiroudelliott/ALongWalk.git
cd ALongWalk
```
Sinon, **téléchargez et extrayez** l’archive du projet manuellement.

---

### **2. Exécuter le programme**
Dans le dossier du projet, utilisez la commande :
```bash
go run main.go final.go
```
Le programme vous demandera de fournir le chemin du fichier d'entrée contenant la grille du labyrinthe.

---

## **Lancer les Tests et Benchmarks**

### **1. Installer les dépendances**
Si ce n'est pas déjà fait, initialisez le module Go et installez **testify** :
```bash
go mod init github.com/mon_utilisateur/ALongWalk  # (remplacez par votre nom d'utilisateur GitHub)
go get github.com/stretchr/testify
go mod tidy
```


### **2. Exécuter les benchmarks**
```bash
go test -bench=. -v
```

# Expliquation du problème et de sa résolution

# Expliquation du problème 

Le problème consiste à trouver le chemin le plus long dans un labyrinthe représenté par une grille. Le labyrinthe est composé de murs (#) et de chemins (.). L'entrée est située sur la première ligne de la grille, et la sortie sur la dernière ligne. L'objectif est de parcourir le labyrinthe et de trouver le chemin qui demande le plus de pas pour aller de l'entrée à la sortie.

# Résolution

Quand le programme commence, il lit le fichier qui contient la grille du labyrinthe. Chaque ligne du fichier est transformée en une ligne de la grille en mémoire. Ensuite, il commence à construire le graphe qui représente les jonctions et les chemins entre elles.

Pour construire ce graphe, le programme parcourt chaque case de la grille et détermine si c'est une jonction. Une jonction est un point où il y a plus de deux chemins possibles, ou bien c'est l'entrée ou la sortie du labyrinthe. Une fois que toutes les jonctions sont identifiées, le programme explore les chemins entre elles. Pour chaque jonction, il suit les chemins autour d'elle jusqu'à ce qu'il trouve une autre jonction. La distance entre les deux jonctions est enregistrée dans le graphe. Ce graphe est non orienté, ce qui signifie que si la jonction A est connectée à la jonction B avec une certaine distance, la jonction B est aussi connectée à la jonction A avec la même distance.

Une fois le graphe construit, le programme le transforme en une version indexée pour faciliter son utilisation. Chaque jonction reçoit un numéro unique, et les connexions entre les jonctions sont stockées dans une liste. Cette liste est triée pour que les chemins les plus longs soient explorés en premier, ce qui aide à trouver rapidement une bonne solution.

Ensuite, le programme détermine où se trouvent l'entrée et la sortie du labyrinthe. Il parcourt la première ligne de la grille pour trouver l'entrée et la dernière ligne pour trouver la sortie.

Le programme commence alors à explorer le graphe pour trouver le chemin le plus long. Il utilise une méthode appelée DFS (exploration en profondeur) qui explore tous les chemins possibles à partir de l'entrée. Pour éviter de perdre du temps, le programme utilise une technique appelée élagage. L'élagage permet d'arrêter l'exploration d'un chemin si on sait qu'il ne pourra pas battre le meilleur chemin déjà trouvé.

Pendant l'exploration, le programme garde une trace des jonctions déjà visitées pour éviter de tourner en rond. Il met à jour la meilleure solution trouvée au fur et à mesure qu'il explore de nouveaux chemins. Une fois que tous les chemins possibles ont été explorés, le programme affiche la longueur du chemin le plus long qu'il a trouvé.