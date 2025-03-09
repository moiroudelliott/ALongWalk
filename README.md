# Installation de Go et exécution du programme

## Installation de Go

Pour exécuter ce programme, vous devez d'abord installer Go (Golang) sur votre machine. Voici les étapes pour installer Go :

1. **Téléchargez Go** :
   - Rendez-vous sur le site officiel de Go : [https://golang.org/dl/](https://golang.org/dl/).
   - Téléchargez la version appropriée pour votre système d'exploitation (Windows, macOS, Linux).

2. **Installez Go** :
   - **Sur Windows** : Exécutez le fichier d'installation téléchargé et suivez les instructions de l'assistant d'installation.
   - **Sur macOS** : Ouvrez le fichier `.pkg` téléchargé et suivez les instructions de l'assistant d'installation.
   - **Sur Linux** : Extrayez l'archive téléchargée dans `/usr/local` ou un autre répertoire de votre choix, puis ajoutez le chemin vers le binaire de Go à votre variable d'environnement `PATH`.

3. **Vérifiez l'installation** :
   - Ouvrez un terminal ou une invite de commandes.
   - Tapez la commande suivante pour vérifier que Go est correctement installé :
     ```bash
     go version
     ```
   - Vous devriez voir la version de Go installée s'afficher.

## Exécution du programme

Une fois Go installé, vous pouvez exécuter le programme en suivant ces étapes :

1. **Clonez ou téléchargez le dépôt** :
   - Si vous utilisez Git, clonez le dépôt contenant le code source :
     ```bash
     git clone <URL_DU_DEPOT>
     ```
   - Sinon, téléchargez et extrayez l'archive du dépôt.

2. **Naviguez vers le répertoire du projet** :
   - Ouvrez un terminal ou une invite de commandes.
   - Naviguez vers le répertoire où se trouve le fichier `main.go` :
     ```bash
     cd chemin/vers/le/repertoire
     ```

3. **Exécutez le programme** :
   - Pour exécuter le programme, utilisez la commande suivante :
     ```bash
     go run main.go
     ```
   - Le programme vous demandera de fournir le chemin vers le fichier d'entrée du labyrinthe. Entrez le chemin complet ou relatif vers le fichier contenant la grille du labyrinthe.

## Exemple d'utilisation

Supposons que vous avez un fichier `labyrinthe.txt` dans le même répertoire que `main.go`. Vous pouvez exécuter le programme comme suit :

```bash
go run main.go
```

Ensuite, entrez le chemin vers le fichier `labyrinthe.txt` lorsque le programme vous le demande :

```
Entrez le chemin vers le fichier d'entrée du labyrinthe : labyrinthe.txt
```

Le programme lira le fichier, trouvera le chemin le plus long dans le labyrinthe, et affichera le résultat.

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