import '/core/data/models/exercise_info_dto.dart';
import '/core/data/models/workout_dto.dart';

abstract class LocalWorkoutDatasource {
  Future<void> saveWorkout(WorkoutDTO workout);
  Future<void> deleteWorkout(int id);
  Future<void> updateWorkout(WorkoutDTO workout);
  Future<WorkoutDTO?> getWorkoutById(int id);
  Future<List<WorkoutDTO>> getAllWorkouts();
  Future<List<ExerciseInfoDTO>> loadInfos();
  Future<void> markInfoFavorite();
  Future<void> unmarkInfoFavorite();
}

class SqfliteDatabase implements LocalWorkoutDatasource {
  @override
  Future<void> deleteWorkout(int id) {
    // TODO: implement deleteWorkout
    throw UnimplementedError();
  }

  @override
  Future<List<WorkoutDTO>> getAllWorkouts() {
    // TODO: implement getAllWorkouts
    throw UnimplementedError();
  }

  @override
  Future<WorkoutDTO?> getWorkoutById(int id) {
    // TODO: implement getWorkoutById
    throw UnimplementedError();
  }

  @override
  Future<List<ExerciseInfoDTO>> loadInfos() {
    // TODO: implement loadInfos
    throw UnimplementedError();
  }  

  @override
  Future<void> markInfoFavorite() {
    // TODO: implement markInfoFavorite
    throw UnimplementedError();
  }

  @override
  Future<void> saveWorkout(WorkoutDTO workout) {
    // TODO: implement saveWorkout
    throw UnimplementedError();
  }

  @override
  Future<void> unmarkInfoFavorite() {
    // TODO: implement unmarkInfoFavorite
    throw UnimplementedError();
  }

  @override
  Future<void> updateWorkout(WorkoutDTO workout) {
    // TODO: implement updateWorkout
    throw UnimplementedError();
  }

}
