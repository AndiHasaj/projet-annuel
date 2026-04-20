<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU | Espace Conseils</title>
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
                <h1 class="font-montserrat font-bold text-ucu-foret text-sm uppercase tracking-widest">Savoir & Tutoriels</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="particulier_dashboard.php" class="text-gray-400 hover:text-ucu-foret">Dashboard</a>
                <a href="particulier_catalogue.php" class="text-gray-400 hover:text-ucu-foret">Catalogue</a>
                <a href="particulier_conseils.php" class="text-ucu-foret border-b-2 border-ucu-lagon pb-1">Conseils</a>
                <a href="particulier_annonces.php" class="text-gray-400 hover:text-ucu-foret">Annonces</a>
                <a href="particulier_conteneurs.php" class="text-gray-400 hover:text-ucu-foret">Conteneurs</a>
                <a href="particulier_planning.php" class="text-gray-400 hover:text-ucu-foret">Planning</a>
                <a href="particulier_profil.php" class="text-gray-400 hover:text-ucu-foret">Profil</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="mb-12 text-center max-w-2xl mx-auto">
            <h2 class="font-montserrat font-bold text-3xl text-ucu-foret uppercase tracking-tighter">Apprendre l'<span class="text-ucu-lagon italic">Upcycling</span></h2>
            <p class="text-sm text-gray-400 mt-2 mb-8">Découvrez les techniques de nos artisans pour transformer vos objets.</p>
            <div class="relative">
                <input type="text" placeholder="Rechercher un tuto (ex: bois, métal, textile...)" class="w-full bg-white p-5 rounded-ucu shadow-sm border border-gray-100 outline-none focus:ring-2 focus:ring-ucu-lagon/30 text-sm italic">
                <button class="absolute right-4 top-3 bg-ucu-foret text-white px-5 py-2 rounded-ucu text-[9px] font-bold uppercase tracking-widest">Chercher</button>
            </div>
        </div>

        <div class="grid grid-cols-1 lg:grid-cols-4 gap-8">
            
            <aside class="lg:col-span-1 space-y-4">
                <h4 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest mb-6">Thématiques</h4>
                <button class="w-full text-left p-4 bg-white rounded-ucu border-l-4 border-ucu-lagon shadow-sm text-xs font-bold text-ucu-foret">🛠️ Travail du Bois</button>
                <button class="w-full text-left p-4 bg-white/50 rounded-ucu hover:bg-white transition-all text-xs font-bold text-gray-400">🧵 Couture & Textile</button>
                <button class="w-full text-left p-4 bg-white/50 rounded-ucu hover:bg-white transition-all text-xs font-bold text-gray-400">🔥 Métallurgie</button>
                <button class="w-full text-left p-4 bg-white/50 rounded-ucu hover:bg-white transition-all text-xs font-bold text-gray-400">🎨 Finitions & Peinture</button>
            </aside>

            <div class="lg:col-span-3 grid grid-cols-1 md:grid-cols-2 gap-8">
                
                <article class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden hover:border-ucu-lagon transition-all">
                    <div class="h-40 bg-gray-200 flex items-center justify-center text-4xl">🪵</div>
                    <div class="p-6">
                        <div class="flex items-center gap-2 mb-3">
                            <span class="bg-ucu-grey text-ucu-lagon text-[8px] font-bold px-2 py-0.5 rounded uppercase">Débutant</span>
                            <span class="text-[8px] text-gray-400 uppercase font-bold tracking-widest">5 min de lecture</span>
                        </div>
                        <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mb-3">Poncer sans abîmer : Les bases</h3>
                        <p class="text-[11px] text-gray-500 leading-relaxed mb-6 italic">Apprenez à choisir le bon grain de papier de verre selon l'essence de votre bois.</p>
                        <a href="#" class="text-ucu-foret font-bold text-[10px] uppercase border-b-2 border-ucu-foret pb-1 hover:text-ucu-lagon hover:border-ucu-lagon transition-all">Lire le guide →</a>
                    </div>
                </article>

                <article class="bg-white rounded-ucu shadow-sm border border-gray-100 overflow-hidden hover:border-ucu-lagon transition-all">
                    <div class="h-40 bg-gray-200 flex items-center justify-center text-4xl">🧴</div>
                    <div class="p-6">
                        <div class="flex items-center gap-2 mb-3">
                            <span class="bg-ucu-grey text-ucu-lagon text-[8px] font-bold px-2 py-0.5 rounded uppercase">Intermédiaire</span>
                            <span class="text-[8px] text-gray-400 uppercase font-bold tracking-widest">12 min de lecture</span>
                        </div>
                        <h3 class="font-montserrat font-bold text-ucu-foret text-sm uppercase mb-3">Recette de vernis naturel</h3>
                        <p class="text-[11px] text-gray-500 leading-relaxed mb-6 italic">Fabriquez votre propre cire de protection à base de cire d'abeille et d'huile de lin.</p>
                        <a href="#" class="text-ucu-foret font-bold text-[10px] uppercase border-b-2 border-ucu-foret pb-1 hover:text-ucu-lagon hover:border-ucu-lagon transition-all">Voir la recette →</a>
                    </div>
                </article>

                <section class="md:col-span-2 bg-ucu-foret p-10 rounded-ucu text-white flex flex-col md:flex-row items-center gap-8 shadow-xl">
                    <div class="flex-1">
                        <h4 class="font-montserrat font-bold text-lg uppercase mb-2 tracking-tighter">Une question technique ?</h4>
                        <p class="text-xs opacity-70 italic leading-relaxed">Posez votre question sur le forum. Nos salariés (animateurs) et artisans vous répondent sous 24h pour vous aider dans votre projet.</p>
                    </div>
                    <button class="bg-ucu-lagon text-white px-8 py-4 rounded-ucu font-bold text-[10px] uppercase tracking-widest hover:bg-white hover:text-ucu-foret transition-all">Accéder au Forum</button>
                </section>
            </div>

        </div>
    </main>

    <footer class="mt-20 py-10 border-t border-gray-100 text-center">
        <p class="text-[9px] font-bold text-gray-300 uppercase tracking-[0.3em]">UpcycleConnect • Espace Conseils • 2026</p>
    </footer>

</body>
</html>