# Parapluie Go

Ce projet fournit une application Go qui interagit avec l'API météorologique Open-Meteo pour récupérer les prévisions de température et de précipitations. Elle calcule la température moyenne et la probabilité moyenne de précipitations, et fournit des recommandations pour savoir quand sortir avec un parapluie ;)

## Fonctionnalités

- Récupération des prévisions de température pour un nombre de jours spécifié
- Calcul de la température moyenne sur la période de prévision
- Récupération des prévisions de probabilité de précipitations pour les 2 prochains jours
- Calcul de la probabilité moyenne de précipitations
- Recommandations de plages horaires pour sortir avec un parapluie (probabilité de précipitations >= 50%)

## Prérequis

- Go (version 1.16 ou supérieure)
- Module `github.com/gin-gonic/gin`
- Module `github.com/go-resty/resty/v2`

## Installation

1. Clonez le dépôt Git :

```
git clone https://github.com/votrecompte/parapluie-go.git
```

2. Accédez au répertoire du projet :

```
cd parapluie-go
```

3. Installez les dépendances :

```
go get ./...
```

## Utilisation

1. Démarrez le serveur Go :

```
go run main.go
```

2. Envoyez des requêtes GET aux endpoints suivants :

- Température moyenne :

```
http://localhost:22222/temperature?latitude=48.8567&longitude=2.3508&previsions_jours=1
```

- Prévisions de précipitations :

```
http://localhost:22222/parapluie?latitude=48.8567&longitude=2.3508
```

Remplacez les valeurs de `latitude` et `longitude` par vos coordonnées géographiques souhaitées.

Vous pouvez modifier le port par défaut en modifiant la ligne 'const PORT = "22222"' dans le fichier main.go

## Tests

Des tests unitaires sont fournis pour vérifier le bon fonctionnement des fonctions principales. Pour exécuter les tests, utilisez la commande suivante :

```
go test ./...
```

Vous devez obtenir cela : 

```
ok  	exercice_meteo	0.171s
ok  	exercice_meteo/meteo	0.244s
```

## Déploiement

Une instance de cette application est déployée et accessible à ces URLs :

- Température moyenne : `https://parapluie.6ops.net/temperature?latitude=52.52&longitude=13.41&previsions_jours=5`
- Prévisions de précipitations : `https://parapluie.6ops.net/parapluie?latitude=52.52&longitude=13.41`

N'hésitez pas à remplacer les valeurs de `latitude` et `longitude` par celles de votre choix.

