import 'package:gym_genius/core/network/dio_service.dart';

import '/core/data/models/workout_dto.dart';

abstract class RemoteWorkoutDatasource {
  Future<void> saveWorkout(WorkoutDTO workout);
  Future<String> getAIDescription(WorkoutDTO workout);
}

class APIWorkoutDatasource implements RemoteWorkoutDatasource {
  APIWorkoutDatasource(this.client);
  final DioService client;

  @override
  Future<String> getAIDescription(WorkoutDTO workout) {
    // TODO: implement getAIDescription
    throw UnimplementedError();
  }

  @override
  Future<void> saveWorkout(WorkoutDTO workout) async {
    client.post('/workouts', data: workout.toJson());
  }
}
