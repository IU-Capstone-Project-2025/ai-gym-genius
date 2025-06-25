import 'dart:async';

import 'package:flutter/material.dart';

class StopwatchController extends ChangeNotifier {
  final Stopwatch _sw = Stopwatch();
  late final Timer _ticker;

  StopwatchController({Duration tick = const Duration(milliseconds: 30)}) {
    _ticker = Timer.periodic(tick, (_) => notifyListeners());
  }

  Duration get elapsed => _sw.elapsed;
  bool get isRunning   => _sw.isRunning;

  void start() { _sw.start(); }
  void stop()  { _sw.stop();  }
  void reset() { _sw.reset(); notifyListeners(); }

  @override
  void dispose() {
    _ticker.cancel();
    _sw.stop();
    super.dispose();
  }
}


class StopwatchWidget extends StatelessWidget {
  final StopwatchController controller;
  final TextStyle? style;

  const StopwatchWidget({
    super.key,
    required this.controller,
    this.style,
  });

  String _fmt(Duration d) {
    String two(int n) => n.toString().padLeft(2, '0');
    return '${two(d.inHours)}:${two(d.inMinutes.remainder(60))}:'
           '${two(d.inSeconds.remainder(60))}';
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: controller,
      builder: (_, __) => Text(
        _fmt(controller.elapsed),
        style: style ??
            TextStyle(
              color: Theme.of(context).colorScheme.onSecondary,
            ),
      ),
    );
  }
}
