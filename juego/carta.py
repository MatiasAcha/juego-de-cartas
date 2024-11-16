class Carta():
    def __init__(self, name, description):
        self.name = name
        self.description = description

class CartaMonstruo(Carta):
    def __init__(self,nombre, description, atributo, puntos):
        super().__init__(nombre, description)
        self.atributo = atributo
        self.puntos = puntos

    def __str__(self):
        return f"{self.name} ({self.atributo}): {self.description} | Poder: {self.puntos}"