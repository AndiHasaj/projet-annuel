<?php
session_start();

$data = [
    "email" => $_POST["email"],
    "mot_de_passe" => $_POST["password"]
];

$ch = curl_init("http://localhost:8080/login/particulier");

curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    "Content-Type: application/json"
]);
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

$response = curl_exec($ch);
$code = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

$result = json_decode($response, true);

if ($code == 200) {
    $_SESSION["user"] = $result;

    header("Location: ../particulier/particulier_dashboard.php");
} else {
    echo $result["error"];
}
?>