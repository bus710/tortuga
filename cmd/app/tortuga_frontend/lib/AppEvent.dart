abstract class AppEvent {
  int OriginalX;
  int OriginalY;
  int DraggedX;
  int DraggedY;
}

class GestureEvent extends AppEvent {
  int OriginalX;
  int OriginalY;
  int DraggedX;
  int DraggedY;

  GestureEvent (originalX, originalY, draggedX, draggedY){
    this.OriginalX = originalX;
    this.OriginalY = originalY;
    this.DraggedX = draggedX;
    this.DraggedY = draggedY;
  }
}