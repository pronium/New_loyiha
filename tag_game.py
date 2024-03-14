import os
import random 
import sys
from PyQt5.QtGui import *
from PyQt5.QtWidgets import *
from PyQt5.QtCore import *
os.system("cls")

class Tag(QMainWindow):
    __ls_btn = []
    __numbers = set()
    def __init__(self):
        super().__init__()
        
        self.setFixedSize(800,600)
        self.setStyleSheet("""background-color: rgb(234, 193, 0);""")
        self.setWindowTitle("Tag game")
        self.setFont(QFont("Consolas",30))

        self.btn1 = self.make_button("1",self)
        self.btn1.move(50,50)

        self.btn2 = self.make_button("2",self)
        self.btn2.move(150,50)

        self.btn3 = self.make_button("3",self)
        self.btn3.move(250,50)

        self.btn4 = self.make_button("4",self)
        self.btn4.move(350,50)

        self.btn5 = self.make_button("5",self)
        self.btn5.move(50,150)

        self.btn6 = self.make_button("6",self)
        self.btn6.move(150,150)

        self.btn7 = self.make_button("7",self)
        self.btn7.move(250,150)

        self.btn8 = self.make_button("8",self)
        self.btn8.move(350,150)

        self.btn9 = self.make_button("9",self)
        self.btn9.move(50,250)

        self.btn10 = self.make_button("10",self)
        self.btn10.move(150,250)

        self.btn11 = self.make_button("11",self)
        self.btn11.move(250,250)

        self.btn12 = self.make_button("12",self)
        self.btn12.move(350,250)

        self.btn13 = self.make_button("13",self)
        self.btn13.move(50,350)

        self.btn14 = self.make_button("13",self)
        self.btn14.move(150,350)

        self.btn15 = self.make_button("15",self)
        self.btn15.move(250,350)

        self.btn16 = self.make_button("",self)
        self.btn16.move(350,350)

        self.new_game = QPushButton("New Game",self)
        self.new_game.setGeometry(500,350,200,60)
        self.new_game.setFont(QFont("Consolas",20))
        self.new_game.clicked.connect(self.start_game)

        self.__ls_btn = [self.btn1,self.btn2,self.btn3,self.btn4,
                         self.btn5,self.btn6,self.btn7,self.btn8,
                         self.btn9,self.btn10,self.btn11,self.btn12,
                         self.btn13,self.btn14,self.btn15,self.btn16]
        
        self.function = [self.click_btn1,self.click_btn2,self.click_btn3,self.click_btn4,
                         self.click_btn5,self.click_btn6,self.click_btn7,self.click_btn8,
                         self.click_btn9,self.click_btn10,self.click_btn11,self.click_btn12,
                         self.click_btn13,self.click_btn14,self.click_btn15,self.click_btn16]
        self.start_game()

        for x,y in zip(self.__ls_btn,self.function):
            x.clicked.connect(y)
        
        self.tablo = QSpinBox(self)
        self.tablo.setGeometry(500,50,150,50)
        self.tablo.setFont(QFont("Consolas",20))
        self.tablo.setValue(0)
        self.tablo.setEnabled(False)
    def start_game(self):
        Tag.__popitka = 0
        while len(Tag.__numbers) != 15:
            Tag.__numbers.add(random.randint(1,15))
        numbers = list(Tag.__numbers)
        random.shuffle(numbers)
        for x,y in zip(self.__ls_btn,numbers):
            x.setText(str(y))
    
    def check_win(self):
        win = True
        for x in range(len(self.__ls_btn)):
            if self.__ls_btn[x].text() == str(x+1):
                continue
            else:
                win = True
        if win:
            msg = QMessageBox(self)
            msg.setText(f"Try count : {Tag.__popitka}")
            msg.setIcon(QMessageBox.Information)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("You Win!")
            msg.show()

    def make_button(self,x,y):
        btn = QPushButton(x,y)
        btn.resize(100,100)
        btn.setFont(QFont("Colibri",30))
        btn.setStyleSheet("""border: 2px solid rgb(7, 204, 4);
                             color: rgb(7, 204, 4);
                             border-radius: 15px;
                             background-color: rgb(7, 5, 122);""")
        return btn
    
    def click_btn1(self):
        Tag.__popitka += 1
        if self.btn2.text() == '':
            self.btn2.setText(self.btn1.text())
            self.btn1.setText("")
        elif self.btn5.text() == '':
            self.btn5.setText(self.btn1.text())
            self.btn1.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    
    def click_btn2(self):
        Tag.__popitka += 1
        if self.btn1.text() == '':
            self.btn1.setText(self.btn2.text())
            self.btn2.setText("")
        elif self.btn3.text() == '':
            self.btn3.setText(self.btn2.text())
            self.btn2.setText('')
        elif self.btn6.text() == '':
            self.btn6.setText(self.btn2.text())
            self.btn2.setText("")
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    def click_btn3(self):
        Tag.__popitka += 1
        if self.btn2.text() == '':
            self.btn2.setText(self.btn3.text())
            self.btn3.setText("")
        elif self.btn4.text() == '':
            self.btn4.setText(self.btn3.text())
            self.btn3.setText('')
        elif self.btn7.text() == '':
            self.btn7.setText(self.btn3.text())
            self.btn3.setText("")
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    def click_btn4(self):
        Tag.__popitka += 1
        if self.btn3.text() == '':
            self.btn3.setText(self.btn4.text())
            self.btn4.setText("")
        elif self.btn8.text() == '':
            self.btn8.setText(self.btn4.text())
            self.btn4.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn5(self):
        Tag.__popitka += 1
        if self.btn1.text() == '':
            self.btn1.setText(self.btn5.text())
            self.btn5.setText("")
        elif self.btn6.text() == '':
            self.btn6.setText(self.btn5.text())
            self.btn5.setText('')
        elif self.btn9.text() == '':
            self.btn9.setText(self.btn5.text())
            self.btn5.setText("")
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    def click_btn6(self):
        Tag.__popitka += 1
        if self.btn7.text() == '':
            self.btn7.setText(self.btn6.text())
            self.btn6.setText("")
        elif self.btn5.text() == '':
            self.btn5.setText(self.btn6.text())
            self.btn6.setText('')
        elif self.btn2.text() == '':
            self.btn2.setText(self.btn6.text())
            self.btn6.setText('')
        elif self.btn10.text() == '':
            self.btn10.setText(self.btn6.text())
            self.btn6.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn7(self):
        Tag.__popitka += 1
        if self.btn8.text() == '':
            self.btn8.setText(self.btn7.text())
            self.btn7.setText("")
        elif self.btn3.text() == '':
            self.btn3.setText(self.btn7.text())
            self.btn7.setText('')
        elif self.btn6.text() == '':
            self.btn6.setText(self.btn7.text())
            self.btn7.setText('')
        elif self.btn11.text() == '':
            self.btn11.setText(self.btn7.text())
            self.btn7.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn8(self):
        Tag.__popitka += 1
        if self.btn4.text() == '':
            self.btn4.setText(self.btn8.text())
            self.btn8.setText("")
        elif self.btn7.text() == '':
            self.btn7.setText(self.btn8.text())
            self.btn8.setText('')
        elif self.btn12.text() == '':
            self.btn12.setText(self.btn8.text())
            self.btn8.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn9(self):
        Tag.__popitka += 1
        if self.btn10.text() == '':
            self.btn10.setText(self.btn9.text())
            self.btn9.setText("")
        elif self.btn13.text() == '':
            self.btn13.setText(self.btn9.text())
            self.btn9.setText('')
        elif self.btn5.text() == '':
            self.btn5.setText(self.btn9.text())
            self.btn9.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn10(self):
        Tag.__popitka += 1
        if self.btn9.text() == '':
            self.btn9.setText(self.btn10.text())
            self.btn10.setText("")
        elif self.btn6.text() == '':
            self.btn6.setText(self.btn10.text())
            self.btn10.setText('')
        elif self.btn11.text() == '':
            self.btn11.setText(self.btn10.text())
            self.btn10.setText('')
        elif self.btn14.text() == '':
            self.btn14.setText(self.btn10.text())
            self.btn10.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn11(self):
        Tag.__popitka += 1
        if self.btn12.text() == '':
            self.btn12.setText(self.btn11.text())
            self.btn11.setText("")
        elif self.btn10.text() == '':
            self.btn10.setText(self.btn11.text())
            self.btn11.setText('')
        elif self.btn15.text() == '':
            self.btn15.setText(self.btn11.text())
            self.btn11.setText('')
        elif self.btn7.text() == '':
            self.btn7.setText(self.btn11.text())
            self.btn11.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn12(self):
        Tag.__popitka += 1
        if self.btn11.text() == '':
            self.btn11.setText(self.btn12.text())
            self.btn12.setText("")
        elif self.btn8.text() == '':
            self.btn8.setText(self.btn12.text())
            self.btn12.setText('')
        elif self.btn16.text() == '':
            self.btn16.setText(self.btn12.text())
            self.btn12.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn13(self):
        Tag.__popitka += 1
        if self.btn14.text() == '':
            self.btn14.setText(self.btn13.text())
            self.btn13.setText("")
        elif self.btn9.text() == '':
            self.btn9.setText(self.btn13.text())
            self.btn13.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
    def click_btn14(self):
        Tag.__popitka += 1
        if self.btn15.text() == '':
            self.btn15.setText(self.btn14.text())
            self.btn14.setText("")
        elif self.btn10.text() == '':
            self.btn10.setText(self.btn14.text())
            self.btn14.setText('')
        elif self.btn13.text() == '':
            self.btn13.setText(self.btn14.text())
            self.btn14.setText("")
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    def click_btn15(self):
        Tag.__popitka += 1
        if self.btn14.text() == '':
            self.btn14.setText(self.btn15.text())
            self.btn15.setText("")
        elif self.btn11.text() == '':
            self.btn11.setText(self.btn15.text())
            self.btn15.setText('')
        elif self.btn16.text() == '':
            self.btn16.setText(self.btn15.text())
            self.btn15.setText("")
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
        self.tablo.setValue(Tag.__popitka)
    def click_btn16(self):
        Tag.__popitka += 1
        if self.btn15.text() == '':
            self.btn15.setText(self.btn16.text())
            self.btn16.setText("")
        elif self.btn12.text() == '':
            self.btn12.setText(self.btn16.text())
            self.btn16.setText('')
        else:
            Tag.__popitka -= 1
            msg = QMessageBox(self)
            msg.setText("Cannot move")
            msg.setIcon(QMessageBox.Critical)
            msg.setFont(QFont("Colibri",18))
            msg.setWindowTitle("E R R O R")
            msg.show()
app = QApplication(sys.argv)
project = Tag()
project.show()
sys.exit(app.exec_())