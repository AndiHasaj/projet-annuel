<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Salarié | Gestion des Événements</title>
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

    <nav class="bg-ucu-lagon border-b border-white/10 sticky top-0 z-50 text-white">
        <div class="max-w-7xl mx-auto px-6 h-20 flex items-center justify-between">
            <div class="flex items-center gap-3">
                <div class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-lagon font-montserrat font-bold italic">SAL</div>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest text-ucu-foret">Espace Salarié</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest text-ucu-foret">
                <a href="salarie_dashboard.php" class="hover:text-white transition-colors">Dashboard</a>
                <a href="salarie_evenements.php" class="border-b-2 border-ucu-foret pb-1">Événements</a>
                <a href="salarie_moderation.php" class="hover:text-white transition-colors">Forums</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-12 gap-12">
            
            <section class="lg:col-span-7 bg-white p-10 rounded-ucu shadow-sm border border-gray-100">
                <header class="mb-10">
                    <h2 class="font-montserrat font-bold text-2xl text-ucu-foret uppercase tracking-tighter">Créer une <span class="text-ucu-lagon italic">Activité</span></h2>
                    <p class="text-[11px] text-gray-400 mt-2 font-bold uppercase tracking-widest">⚠️ Soumis à validation du responsable</p>
                </header>

                <form class="space-y-6">
                    <div class="grid grid-cols-2 gap-6">
                        <div class="col-span-2">
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Titre de l'événement</label>
                            <input type="text" placeholder="ex: Atelier Soudure Initiation" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none focus:ring-2 focus:ring-ucu-lagon/20">
                        </div>
                        
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Type</label>
                            <select class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                                <option>Formation</option>
                                <option>Atelier Collaboratif</option>
                                <option>Événement News</option>
                            </select>
                        </div>

                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Prix (Stripe)</label>
                            <input type="number" placeholder="0.00 €" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                        </div>
                    </div>

                    <div>
                        <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Description & Objectifs</label>
                        <textarea rows="4" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none" placeholder="Décrivez le contenu pédagogique..."></textarea>
                    </div>

                    <div class="grid grid-cols-2 gap-6">
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Date</label>
                            <input type="date" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                        </div>
                        <div>
                            <label class="text-[9px] font-bold uppercase text-gray-400 mb-2 block">Lieu / Salle</label>
                            <input type="text" placeholder="Atelier A1 - Ivry" class="w-full bg-ucu-grey p-4 rounded-ucu text-xs outline-none">
                        </div>
                    </div>

                    <button type="submit" class="w-full bg-ucu-foret text-white py-4 rounded-ucu font-montserrat font-bold text-[10px] uppercase tracking-[0.2em] shadow-lg hover:bg-ucu-lagon transition-all">
                        Envoyer pour validation
                    </button>
                </form>
            </section>

            <section class="lg:col-span-5 space-y-6">
                <div class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest">Mes demandes récentes</h3>
                    
                    <div class="space-y-4">
                        <div class="p-5 border border-gray-100 rounded-ucu flex justify-between items-center group">
                            <div>
                                <p class="text-[11px] font-bold text-ucu-foret uppercase">Tuto : Vernis naturel</p>
                                <p class="text-[9px] text-gray-400 italic mt-1">Soumis le 14/03</p>
                            </div>
                            <span class="text-[8px] bg-orange-100 text-orange-500 font-bold px-2 py-1 rounded uppercase">Attente</span>
                        </div>

                        <div class="p-5 border border-gray-100 rounded-ucu flex justify-between items-center bg-green-50/30">
                            <div>
                                <p class="text-[11px] font-bold text-ucu-foret uppercase">Atelier Bois Flotté</p>
                                <p class="text-[9px] text-gray-400 italic mt-1">Validé par Admin</p>
                            </div>
                            <span class="text-[8px] bg-green-500 text-white font-bold px-2 py-1 rounded uppercase italic">Publié</span>
                        </div>

                        <div class="p-5 border border-gray-100 rounded-ucu flex justify-between items-center">
                            <div>
                                <p class="text-[11px] font-bold text-gray-400 uppercase line-through">Stage Forge</p>
                                <p class="text-[9px] text-red-400 italic mt-1 font-bold">Refusé : Lieu indisponible</p>
                            </div>
                            <button class="text-xs">⚙️</button>
                        </div>
                    </div>
                </div>

                <div class="bg-ucu-foret p-6 rounded-ucu text-white">
                    <p class="text-[10px] font-bold uppercase mb-2 text-ucu-lagon">Note de service</p>
                    <p class="text-[10px] leading-relaxed opacity-80 italic">
                        Les propositions d'ateliers doivent être soumises au minimum 15 jours avant la date prévue pour permettre la génération des supports PDF et l'ouverture des paiements Stripe.
                    </p>
                </div>
            </section>

        </div>
    </main>

</body>
</html>