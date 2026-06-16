<?php

$nom = $_POST["nom"];
$prenom = $_POST["prenom"];
$email = $_POST["email"];
$password = $_POST["password"];

$data = [
    "nom" => $nom,
    "prenom" => $prenom,
    "email" => $email,
    "mot_de_passe" => $password
];

$ch = curl_init("http://localhost:8080/particuliers/");

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    "Content-Type: application/json"
]);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

$response = curl_exec($ch);
if ($response === false) {
    echo "Erreur cURL : " . curl_error($ch);
}
$httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($httpCode == 200 || $httpCode == 201) {
    header("Location: ../particulier/particulier_dashboard.php");
    exit;
} else {
    echo "Erreur inscription : " . $response;
}
?>