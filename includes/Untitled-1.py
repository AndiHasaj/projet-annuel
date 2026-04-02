# --- Partie 1 : Afficher les nombres de 1 à 10 ---
print("Nombres de 1 à 10 :")
for i in range(1, 11): # range(début, fin_exclue)
    print(i)

# --- Partie 2 : Somme des nombres pairs ---
# Demander un entier à l'utilisateur
n_saisi = int(input("\nSaisissez un nombre entier : "))

somme = 0
for i in range(2, n_saisi + 1):
    if i % 2 == 0: # Si le reste de la division par 2 est 0, c'est pair
        somme += i

print(f"La somme des nombres pairs entre 2 et {n_saisi} est : {somme}")