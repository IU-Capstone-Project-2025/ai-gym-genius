import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class StartWorkoutPage extends StatelessWidget {
  const StartWorkoutPage({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: CupertinoButton.tinted(
          child: Text(
            'Start a workout',
          ),
          onPressed: () {
            Navigator.pushNamed(context, '/training_process');
          },
        ),
      ),
    );
  }
}
