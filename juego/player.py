from juego import mazo

class Player():
    def __init__(self, name, deck):
        self.name = name
        self.deck = deck
        self.cementerio = mazo.Mazo()
        self.vida = 20
        self.mano = []

    def robar(self):
        carta = self.deck.robar()
        if carta == None:
            return True
        else:
            self.mano.append(carta)
            return False
        
    def iniciarturno(self):
        pass
    
    def getmano(self):
        # Mostrar mano del jugador
        return '\n'.join(str(carta) for carta in self.mano) if self.mano else "No tiene cartas en la mano."
