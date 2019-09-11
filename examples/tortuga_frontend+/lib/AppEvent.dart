abstract class AppEvent {
  String buttonName;
}

class ButtonEvent extends AppEvent {
  String buttonName;
  ButtonEvent(String buttonName){
    this.buttonName = buttonName;
  }
}