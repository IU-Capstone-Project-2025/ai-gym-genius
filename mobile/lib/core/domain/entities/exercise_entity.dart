import 'exercise_info_entity.dart';
import 'exercise_set_entity.dart';

class ExerciseEntity {
  final ExerciseInfoEntity exerciseInfo;
  final List<ExerciseSetEntity> sets;

  ExerciseEntity({required this.exerciseInfo, required this.sets});
}