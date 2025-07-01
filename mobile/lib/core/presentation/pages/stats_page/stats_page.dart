import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gym_genius/core/domain/entities/workout_entity.dart';
import 'package:gym_genius/core/domain/repositories/workout_repository.dart';
import 'package:gym_genius/di.dart';

import 'stats_widgets.dart';

class StatsPage extends StatefulWidget {
  const StatsPage({super.key});

  @override
  State<StatsPage> createState() => _StatsPageState();
}

class _StatsPageState extends State<StatsPage> {
  late final Future<List<WorkoutEntity>> _workoutsF;

  @override
  void initState() {
    super.initState();
    _workoutsF = getIt<WorkoutRepository>().fetchWorkouts();
  }

  @override
  Widget build(BuildContext context) {
    final schema = Theme.of(context).colorScheme;

    return Scaffold(
      appBar: CupertinoNavigationBar.large(
        border: null,
        backgroundColor: schema.surfaceContainer,
        largeTitle: Text(
          'Statistics',
          style: TextStyle().copyWith(color: schema.primary),
        ),
      ),
      body: FutureBuilder<List<WorkoutEntity>>(
        future: _workoutsF,
        builder: (_, snap) {
          if (snap.connectionState == ConnectionState.waiting) {
            return const Center(child: CircularProgressIndicator.adaptive());
          }
          if (snap.hasError) {
            return Center(child: Text('Error: ${snap.error}'));
          }

          final workouts = snap.data!;
          // normalise each startTime to a pure date
          final workoutDays = workouts
              .map((w) => DateTime(
                  w.startTime.year, w.startTime.month, w.startTime.day))
              .toSet();

          return SingleChildScrollView(
              child: ActivityGrid(workoutDays: workoutDays));
        },
      ),
    );
  }
}
