<?php
// Récupération des données Marketing via l'API Go
$data = @file_get_contents("http://localhost:8080/api/admin/marketing");
$marketing = json_decode($data, true);

// Fallback si le serveur Go est éteint
if (!$marketing) {
    $marketing = [
        "boosts_count" => 0, 
        "active_boosts" => []
    ];
}
?>
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Admin | Marketing & Notifications</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;600&family=Montserrat:wght@700&display=swap" rel="stylesheet">
    <script>
        tailwind.config = {
            theme: {
                extend: {
                    colors: {
                        'ucu-foret': '#1E4E4E',
                        'ucu-lagon': '#6A9A9A',
                        'ucu-grey': '#F8FAFC',
                    },
                    fontFamily: {
                        'sans': ['Inter', 'sans-serif'],
                        'montserrat': ['Montserrat', 'sans-serif'],
                    },
                    borderRadius: { 'ucu': '8px' }
                }
            }
        }
    </script>
</head>
<body class="bg-ucu-grey font-sans text-gray-800">

    <?php include '../includes/sidebar.php'; ?>

    <main class="lg:ml-72 p-12">
        <header class="mb-12">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Communication <span class="text-ucu-lagon italic">& Offres</span></h2>
            <p class="text-sm text-gray-400 mt-2">Diffusez des notifications massives et gérez la visibilité du catalogue via OneSignal.</p>
        </header>

        <div class="grid grid-cols-1 lg:grid-cols-12 gap-10">
            
            <section class="lg:col-span-7 space-y-8">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <div class="flex items-center gap-3 mb-8">
                        <span class="text-2xl">📡</span>
                        <h3 class="font-montserrat font-bold text-xs uppercase text-ucu-foret tracking-widest">Nouvelle Notification Push</h3>
                    </div>

                    <form action="#" method="POST" class="space-y-6">
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Segment d'audience</label>
                            <select name="audience" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:ring-2 focus:ring-ucu-lagon/20 cursor-pointer">
                                <option value="all">Tous les utilisateurs</option>
                                <option value="particuliers">Particuliers uniquement</option>
                                <option value="pros">Professionnels uniquement</option>
                            </select>
                        </div>

                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Titre du message</label>
                            <input type="text" name="title" required placeholder="ex: Nouvelle formation disponible !" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none font-bold focus:bg-white focus:ring-1 focus:ring-ucu-lagon/30 transition-all">
                        </div>

                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Contenu (Max 150 caractères)</label>
                            <textarea name="message" rows="3" required class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:bg-white focus:ring-1 focus:ring-ucu-lagon/30 transition-all" placeholder="Décrivez l'offre ou l'alerte..."></textarea>
                        </div>

                        <div class="pt-4 border-t border-gray-50 flex items-center justify-between">
                            <label class="flex items-center gap-2 cursor-pointer group">
                                <input type="checkbox" class="accent-ucu-lagon w-4 h-4">
                                <span class="text-[10px] text-gray-400 font-bold uppercase group-hover:text-ucu-foret transition-colors">Programmer l'envoi</span>
                            </label>
                            <button type="submit" class="bg-ucu-foret text-white px-8 py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-widest shadow-lg hover:bg-ucu-lagon transform hover:-translate-y-1 transition-all">
                                Envoyer via OneSignal
                            </button>
                        </div>
                    </form>
                </div>
            </section>

            <aside class="lg:col-span-5 space-y-8">
                <div class="bg-ucu-foret p-8 rounded-ucu text-white relative overflow-hidden shadow-xl border-l-4 border-ucu-lagon">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-lagon mb-4 tracking-widest">Visibilité payante (Stripe)</h3>
                    <p class="text-2xl font-montserrat font-bold mb-1"><?= $marketing['boosts_count'] ?> <span class="text-[10px] font-normal opacity-60">Boosts actifs</span></p>
                    <p class="text-[9px] opacity-70 italic mb-6 leading-relaxed">Liste des artisans ayant payé pour une mise en avant temporaire.</p>
                    
                    <div class="space-y-3">
                        <?php if(!empty($marketing['active_boosts'])): ?>
                            <?php foreach($marketing['active_boosts'] as $boost): ?>
                            <div class="bg-white/10 p-3 rounded flex justify-between items-center text-[10px] border border-white/5 hover:bg-white/20 transition-all">
                                <span><?= $boost['nom'] ?></span>
                                <span class="font-bold text-ucu-lagon">Fin : <?= $boost['expiration'] ?></span>
                            </div>
                            <?php endforeach; ?>
                        <?php else: ?>
                            <p class="text-[10px] opacity-50 italic">Aucun boost actif actuellement.</p>
                        <?php endif; ?>
                    </div>
                </div>

                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Geofencing</h3>
                    <p class="text-[10px] text-gray-500 leading-relaxed mb-6">
                        Envoyez une notification automatique aux particuliers situés à moins de 5km d'un conteneur qui vient d'être vidé.
                    </p>
                    <button class="w-full border-2 border-ucu-lagon text-ucu-lagon py-3 rounded-ucu font-bold text-[9px] uppercase tracking-widest hover:bg-ucu-lagon hover:text-white transition-all">
                        Activer le Geofencing
                    </button>
                </div>
            </aside>

        </div>
    </main>
</body>
</html>