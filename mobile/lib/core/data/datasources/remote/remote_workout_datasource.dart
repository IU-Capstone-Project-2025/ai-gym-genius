import '/core/data/models/workout_dto.dart';

abstract class RemoteWorkoutDatasource {
  Future<void> saveWorkout(WorkoutDTO workout);
  Future<String> getAIDescription(WorkoutDTO workout);
}

class APIWorkoutDatasource implements RemoteWorkoutDatasource {
  @override
  Future<String> getAIDescription(WorkoutDTO workout) {
    // TODO: implement getAIDescription
    throw UnimplementedError();
  }

  @override
  Future<void> saveWorkout(WorkoutDTO workout) {
    // TODO: implement saveWorkout
    throw UnimplementedError();
  }

}