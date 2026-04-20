# Schéma de la base de données

```mermaid
erDiagram

PARTICULIER ||--o{ ANNONCE : cree
PARTICULIER ||--o{ PAIEMENT : effectue
PARTICULIER ||--o{ OBJET : depose
ENTREPRISE ||--o{ PAIEMENT : effectue
ENTREPRISE ||--o{ ANNONCE : cree
CONTENEUR ||--o{ OBJET : contient
ADMIN ||--o{ PARTICULIER : gere
ADMIN ||--o{ ANNONCE : gere
ADMIN ||--o{ CONTENEUR : gere
ADMIN ||--o{ PAIEMENT : gere
ADMIN ||--o{ ENTREPRISE : gere
ADMIN ||--o{ SERVICES : gere
ADMIN ||--o{ SUIVI_FINANCIER : gere
ADMIN ||--o{ SALARIÉ : gere
ADMIN ||--o{ OBJET : valide
ENTREPRISE ||--o{ OBJET : recupere



PARTICULIER {
  int id_particulier PK
  string nom
  string prenom
  string email
  string mot_de_passe
  int score(0)
  date derniere_connexion
}

SALARIÉ {
  int id_salarie PK
  string nom
  string prenom
  string email_pro
  string mot_de_passe
  string fonction
  string type_contrat
  string salaire
  date date_debut_contrat
  date derniere_connexion
}

ENTREPRISE {
  int id_entreprise PK
  string nom
  string email
  string mot_de_passe
  string(14) numero_siret
  string(20) telephone
  string site_web
  string adresse
  string(5) code_postal
  string ville
}

ANNONCE {
  int id_annonce PK
  string titre
  string description
  decimal prix
  date date_publication
  int id_particulier FK
  int id_entreprise FK
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
  string matricule
  string localisation
  string statut
  string occupation
  string code
}

SERVICES {
  int id_service PK
  string type_service
  string titre
  decimal prix
  string description
  date date
}

PAIEMENT {
  int id_paiement PK
  float montant
  date date_paiement
  string type_paiement
  string statut
  int id_payeur FK
}

SUIVI_FINANCIER {
  decimal chiffre_affaire
  decimal charges
  decimal resultat_net

}

ADMIN {
  int id_admin PK
  string mot_de_passe
}
```