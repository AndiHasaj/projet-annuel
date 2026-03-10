# Schéma de la base de données

```mermaid
erDiagram

USER ||--o{ ANNONCE : cree
USER ||--o{ PAIEMENT : effectue
CONTENEUR ||--o{ OBJET : contient

USER {
  int id_utilisateur PK
  string nom
  string prenom
  string email
  string mot_de_passe
  string role
}

ANNONCE {
  int id_annonce PK
  string titre
  string description
  float prix
  date date_publication
  int id_utilisateur FK
}

OBJET {
  int id_objet PK
  string nom
  string type_objet
  string localisation
  int id_conteneur FK
}

CONTENEUR {
  int id_conteneur PK
  string localisation
  string statut
  string code
}

SERVICES {
  int id_service PK
  string type_service
  string titre
  float prix
  string description
  date date
}

PAIEMENT {
  int id_paiement PK
  float montant
  date date_paiement
  string type_paiement
  string statut
  int id_utilisateur FK
}
```
