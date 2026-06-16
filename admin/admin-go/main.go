package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DashboardStats struct {
	Revenus24h        int `json:"revenus_24h"`
	NouveauxPros      int `json:"nouveaux_pros"`
	SignalementsBox   int `json:"signalements_box"`
	FormationsAttente int `json:"formations_attente"`
	TotalMensuel      int `json:"total_mensuel"`
	Croissance        int `json:"croissance"`
	DataAbonnements   int `json:"data_abonnements"`
	DataFormations    int `json:"data_formations"`
	DataBoosts        int `json:"data_boosts"`
}

type Box struct {
	Matricule        string `json:"matricule"`
	Localisation     string `json:"localisation"`
	Ville            string `json:"ville"`
	Occupation       int    `json:"occupation"`
	DerniereCollecte string `json:"derniere_collecte"`
}

type User struct {
	Nom    string `json:"nom"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Status string `json:"status"`
	Detail string `json:"detail"`
}

type Transaction struct {
	Date        string  `json:"date"`
	Designation string  `json:"designation"`
	Type        string  `json:"type"`
	Montant     float64 `json:"montant"`
}

type FinanceData struct {
	CATotal        float64       `json:"ca_total"`
	ChargesTotales float64       `json:"charges_totales"`
	ResultatNet    float64       `json:"resultat_net"`
	TauxBenefice   int           `json:"taux_benefice"`
	Transactions   []Transaction `json:"transactions"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	stats := DashboardStats{
		Revenus24h: 1240, NouveauxPros: 8, SignalementsBox: 2, FormationsAttente: 5,
		TotalMensuel: 14500, Croissance: 12, DataAbonnements: 8000, DataFormations: 4500, DataBoosts: 2000,
	}
	json.NewEncoder(w).Encode(stats)
}

func handleLogistique(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	boxes := []Box{
		{"BOX-IVRY-001", "Centre Commercial Quays", "94200 Ivry", 90, "14/03/2026"},
		{"BOX-PARIS-042", "Place de la Bastille", "75011 Paris", 30, "12/03/2026"},
		{"BOX-LYON-005", "Gare Part-Dieu", "69003 Lyon", 15, "18/03/2026"},
	}
	json.NewEncoder(w).Encode(boxes)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	users := []User{
		{"Atelier du Bois d'Ivry", "contact@boisivry.pro", "Professionnel", "À Valider", "KBIS_2026.pdf"},
		{"LACI Senard", "l.senard@upcycleconnect.fr", "Salarié", "Actif", "Contrat Interne"},
		{"Andi HASAJ", "andi.h@gmail.com", "Particulier", "Actif", "450 pts"},
	}
	json.NewEncoder(w).Encode(users)
}

func handleFinance(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	finance := FinanceData{
		CATotal:        24590.00,
		ChargesTotales: 8120.00,
		ResultatNet:    16460.00,
		TauxBenefice:   14,
		Transactions: []Transaction{
			{"15/03/2026", "Abonnement Annuel - Atelier du Bois", "Revenu", 450.00},
			{"14/03/2026", "Hébergement Cloud (Google Cloud)", "Charge", 120.00},
			{"12/03/2026", "Vente Formation Upcycling Débutant", "Revenu", 85.00},
			{"10/03/2026", "Maintenance Serveur API Go", "Charge", 200.00},
		},
	}
	json.NewEncoder(w).Encode(finance)
}

func main() {
	http.HandleFunc("/api/admin/stats", handleDashboard)
	http.HandleFunc("/api/admin/logistique", handleLogistique)
	http.HandleFunc("/api/admin/users", handleUsers)
	http.HandleFunc("/api/admin/finance", handleFinance)

	fmt.Println("🌐 API UpcycleConnect : http://localhost:8080")
	fmt.Println("🚀 Routes prêtes : stats, logistique, users, finance")
	
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erreur lors du lancement du serveur : %s\n", err)
	}
}