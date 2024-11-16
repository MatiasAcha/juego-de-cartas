import random

class Mazo():
    def __init__(self):
        self.mazo = []
        self.total = 0

    def agregarcarta(self, carta):
        if self.total >= 60:
            print("El mazo está completo. No se puede agregar más cartas.")
        else:
            self.mazo.append(carta)
            self.total += 1

    def robar(self):
        if self.total == 0:
            print("El mazo está vacío.")
            return None
        else:
            carta = self.mazo.pop()
            self.total -= 1
            return carta
    
    def mezclar(self):
        random.shuffle(self.mazo)

    def __str__(self):
        return '\n'.join(str(carta) for carta in self.mazo)