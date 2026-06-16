<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>UCU Staff | Espace Animateur</title>
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
                <div class="w-10 h-10 bg-white rounded-ucu flex items-center justify-center text-ucu-lagon font-montserrat font-bold italic">STAFF</div>
                <h1 class="font-montserrat font-bold text-sm uppercase tracking-widest">Espace Formateur</h1>
            </div>
            <div class="flex gap-6 text-[10px] font-bold uppercase tracking-widest">
                <a href="staff_dashboard.php" class="text-ucu-foret">Dashboard</a>
                <a href="staff_evenements.php" class="hover:text-ucu-foret transition-colors">Événements</a>
                <a href="staff_moderation.php" class="hover:text-ucu-foret transition-colors">Modération</a>
            </div>
        </div>
    </nav>

    <main class="max-w-7xl mx-auto px-6 py-12">
        
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            
            <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret mb-6 tracking-widest italic">Suivi des validations</h3>
                <div class="space-y-4">
                    <div class="p-4 bg-orange-50 rounded-ucu border border-orange-100">
                        <p class="text-[9px] font-bold text-orange-600 uppercase">En attente (2)</p>
                        <p class="text-xs font-medium text-gray-700 mt-1">Atelier : "Soudure Arc Niv.1"</p>
                    </div>
                    <div class="p-4 bg-green-50 rounded-ucu border border-green-100">
                        <p class="text-[9px] font-bold text-green-600 uppercase">Validé (14)</p>
                        <p class="text-xs font-medium text-gray-700 mt-1">News : "Le recyclage du PET"</p>
                    </div>
                </div>
            </section>

            <div class="lg:col-span-2 space-y-8">
                
                <section class="bg-white p-8 rounded-ucu shadow-sm border border-gray-100">
                    <div class="flex justify-between items-center mb-6">
                        <h3 class="font-montserrat font-bold text-[10px] uppercase text-ucu-foret tracking-widest">Forum : Signalements récents</h3>
                        <span class="bg-red-500 text-white text-[8px] font-bold px-2 py-0.5 rounded-full">3 NOUVEAUX</span>
                    </div>
                    <div class="divide-y divide-gray-50">
                        <div class="py-4 flex justify-between items-center">
                            <div>
                                <p class="text-xs font-bold uppercase text-ucu-foret">Sujet : "Vernis bio..."</p>
                                <p class="text-[10px] text-gray-400 italic">Signalé pour : Publicité déguisée</p>
                            </div>
                            <button class="text-[9px] font-bold text-ucu-lagon uppercase">Modérer</button>
                        </div>
                    </div>
                </section>

                <section class="bg-ucu-foret p-10 rounded-ucu shadow-xl text-white relative overflow-hidden">
                    <div class="relative z-10 flex flex-col md:flex-row justify-between items-center">
                        <div>
                            <span class="text-[9px] font-bold text-ucu-lagon uppercase tracking-widest">Prochain Direct</span>
                            <h3 class="font-montserrat font-bold text-2xl uppercase mt-2">Atelier Tapisserie</h3>
                            <p class="text-xs opacity-60 mt-1 italic">Demain à 14h00 • Centre Ivry</p>
                        </div>
                        <div class="bg-white/10 p-4 rounded-ucu text-center border border-white/10">
                            <p class="text-2xl font-montserrat font-bold">18</p>
                            <p class="text-[8px] uppercase font-bold opacity-60">Inscrits</p>
                        </div>
                    </div>
                    <div class="absolute -right-4 -bottom-4 text-7xl opacity-5">🎓</div>
                </section>
            </div>

        </div>
    </main>
</body>
</html>