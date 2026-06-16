<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Catalogue & Services</title>
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

    <nav class="bg-white border-b border-gray-100 sticky top-0 z-50">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <a href="particulier_dashboard.php" class="w-10 h-10 bg-ucu-foret rounded-ucu flex items-center justify-center text-white font-montserrat font-bold italic">UCU</a>
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Boutique & Savoir</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Catalogue</a>
                <a href="particulier_conseils.php" class="text-gray-400 hover:text-ucu-foret">Conseils</a>
                <a href="particulier_annonces.php" class="text-gray-400 hover:text-ucu-foret">Annonces</a>
                <a href="particulier_conteneurs.php" class="text-gray-400 hover:text-ucu-foret">Conteneurs</a>
                <a href="particulier_planning.php" class="text-gray-400 hover:text-ucu-foret">Planning</a>
                <a href="particulier_profil.php" class="text-gray-400 hover:text-ucu-foret">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <header class="mb-12 flex flex-col md:flex-row justify-between items-end gap-6">
            <div>
                <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Catalogue des <span class="text-ucu-lagon italic">Services</span></h2>
                <p class="text-sm text-gray-400 mt-2">Réservez vos formations et événements d'upcycling.</p>
            </div>
            <div class="flex gap-4">
                <select class="bg-white border border-gray-100 px-4 py-3 rounded-ucu text-[10px] font-bold uppercase tracking-widest outline-none focus:ring-2 focus:ring-ucu-lagon/20">
                    <option>Toutes les catégories</option>
                    <option>Formations</option>
                    <option>Événements</option>
                    <option>Services Artisans</option>
                </select>
            </div>
        </header>

        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            
            <article class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden hover:shadow-md transition-all group">
                <div class="h-48 bg-gray-200 relative overflow-hidden">
                    <div class="absolute inset-0 bg-ucu-foret/20 group-hover:bg-ucu-foret/0 transition-all"></div>
                    <div class="absolute top-4 left-4 bg-ucu-foret text-white text-[9px] font-bold px-3 py-1 rounded-full uppercase">Formation</div>
                    <div class="w-full h-full flex items-center justify-center text-4xl bg-ucu-grey">🪑</div>
                </div>
                <div class="p-6">
                    <div class="flex justify-between items-start mb-2">
                        <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase leading-tight">Rénovation de Siège : <br>Techniques de Tapissier</h3>
                        <p class="font-montserrat font-bold text-ucu-lagon">45€</p>
                    </div>
                    <p class="text-[11px] text-gray-500 italic mb-6">Apprenez à redonner vie à vos vieux fauteuils avec l'artisan local d'Ivry.</p>
                    <div class="flex items-center justify-between pt-4 border-t border-gray-50">
                        <span class="text-[9px] font-bold text-gray-400 uppercase tracking-widest">📅 24 Mars 2026</span>
                        <button class="bg-ucu-foret text-white px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-ucu-lagon transition-all">Réserver (Stripe)</button>
                    </div>
                </div>
            </article>

            <article class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden hover:shadow-md transition-all group">
                <div class="h-48 bg-gray-200 relative overflow-hidden">
                    <div class="absolute top-4 left-4 bg-ucu-lagon text-white text-[9px] font-bold px-3 py-1 rounded-full uppercase">Événement</div>
                    <div class="w-full h-full flex items-center justify-center text-4xl bg-ucu-grey">🤝</div>
                </div>
                <div class="p-6">
                    <div class="flex justify-between items-start mb-2">
                        <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase leading-tight">Grande Bourse à <br>l'Upcycling Collectif</h3>
                        <p class="font-montserrat font-bold text-ucu-lagon uppercase">Gratuit</p>
                    </div>
                    <p class="text-[11px] text-gray-500 italic mb-6">Venez échanger vos chutes de matériaux et trouver l'inspiration pour vos projets.</p>
                    <div class="flex items-center justify-between pt-4 border-t border-gray-50">
                        <span class="text-[9px] font-bold text-gray-400 uppercase tracking-widest">📍 Centre Ivry</span>
                        <button class="bg-ucu-lagon text-white px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-ucu-foret transition-all">S'inscrire</button>
                    </div>
                </div>
            </article>

            <article class="bg-ucu-foret rounded-ucu shadow-xl border border-ucu-foret overflow-hidden relative group">
                <div class="p-8 text-white h-full flex flex-col">
                    <div class="mb-6">
                        <span class="text-[9px] font-bold text-ucu-lagon uppercase tracking-[0.2em]">Service Premium</span>
                        <h3 class="font-montserrat font-bold text-xl uppercase mt-2 tracking-tighter">Diagnostic <br>Check à domicile</h3>
                    </div>
                    <p class="text-[11px] opacity-70 italic mb-8 flex-1 leading-relaxed">Un expert UpcycleConnect vient chez vous pour expertiser vos meubles et proposer un plan de revalorisation.</p>
                    <div class="mt-auto flex items-center justify-between border-t border-white/10 pt-6">
                        <p class="font-montserrat font-bold text-ucu-lagon">25€ <span class="text-[8px] text-white opacity-40">/ visite</span></p>
                        <button class="bg-white text-ucu-foret px-4 py-2 rounded-ucu font-bold text-[9px] uppercase hover:bg-ucu-lagon hover:text-white transition-all">Prendre RDV</button>
                    </div>
                    <div class="absolute -right-8 -top-8 w-24 h-24 bg-white/5 rounded-full"></div>
                </div>
            </article>

        </div>

        <section class="mt-16 bg-white p-10 rounded-ucu border border-gray-100 flex flex-col md:flex-row items-center gap-10 shadow-sm">
            <div class="w-24 h-24 bg-ucu-grey rounded-full flex items-center justify-center text-3xl border-2 border-ucu-lagon border-dashed">🎁</div>
            <div class="flex-1">
                <h4 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mb-2">Utilisez votre Upcycling Score</h4>
                <p class="text-[11px] text-gray-400 leading-relaxed italic">Vous avez actuellement <span class="text-ucu-lagon font-bold">750 points</span>. Échangez-les contre des coupons de réduction valables sur toutes les formations de tapisserie et soudure.</p>
            </div>
            <button class="w-full md:w-auto border-2 border-ucu-foret text-ucu-foret px-8 py-4 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-ucu-foret hover:text-white transition-all">
                Convertir mes points
            </button>
        </section>

    </main>

</body>
</html>