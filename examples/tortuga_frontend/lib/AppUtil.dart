import 'package:flutter_web/material.dart';

class ButtonPainter extends CustomPainter {
  bool state = false;

  ButtonPainter(bool state) {
    this.state = state;
  }

  @override
  void paint(Canvas canvas, Size size) {
    Color selectedColor = Colors.white;
    Offset start1 = Offset(1, 1);
    Offset end1 = Offset(1, 40);
    Offset start2 = Offset(1, 1);
    Offset end2 = Offset(40, 1);
    Offset start3 = Offset(1, 40);
    Offset end3 = Offset(40, 1);

    // double radius = 10;
    Paint brush = Paint()
      ..color = selectedColor
      ..style = PaintingStyle.stroke
      ..strokeCap = StrokeCap.round
      ..strokeWidth = 5;

    if (this.state == true) {
      canvas.drawLine(start1, end1, brush);
      canvas.drawLine(start2, end2, brush);
      canvas.drawLine(start3, end3, brush);
    }
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) {
    return true;
  }
}
