class Human:
    def __init__(self, name: str):
        self.name = name

    def say(self, text: str) -> None:
        print(self.name + "< " + text)

    def mr(self) -> str:
        return "Mr."+self.name

    def my_name(self) -> str:
        return "myname"
