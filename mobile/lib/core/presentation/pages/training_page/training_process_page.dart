import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:gym_genius/core/presentation/shared/warnings.dart';
import 'package:gym_genius/di.dart' show getIt;
import 'package:flutter_bloc/flutter_bloc.dart';

import '/core/presentation/bloc/training_bloc.dart';
import '/core/presentation/bloc/training_event.dart';
import '/core/presentation/bloc/training_state.dart';
import '/core/presentation/pages/training_page/widgets/stopwatch_widget.dart';
import 'picking_exercise_page.dart';
import 'widgets/widgets.dart' show StopwatchWidget, ExpandableExerciseTile;

class TrainingProcessPage extends StatefulWidget {
  const TrainingProcessPage({super.key});

  @override
  State<TrainingProcessPage> createState() => _TrainingProcessPageState();
}

class _TrainingProcessPageState extends State<TrainingProcessPage> {
  late StopwatchController _timerCtrl;

  @override
  void initState() {
    super.initState();
    _timerCtrl = StopwatchController()..start();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final double heightOfPinnedContainer =
        MediaQuery.of(context).size.height / 5;

    return BlocProvider(
      create: (context) => getIt<TrainingBloc>(),
      child: BlocListener<TrainingBloc, TrainingState>(
        listener: (context, state) { 
          if (state.submitTrainingStatus == SubmitTrainingStatus.success) {
            Navigator.pop(context);
          }

          else if (state.submitTrainingStatus == SubmitTrainingStatus.failureEmptySets) {
            Warnings.showSubmitIncompleteWarning(context);
          }
        },
        child: BlocBuilder<TrainingBloc, TrainingState>(
              builder: (context, state) {
                final bloc = context.read<TrainingBloc>();
      
                return Scaffold(
                  body: Stack(
                    children: [
                      _buildSlivers(
                        StopwatchWidget(controller: _timerCtrl),
                        bloc.state.exercises.isNotEmpty,
                        heightOfPinnedContainer,
                        bloc,
                      ),
                      _buildPinnedContainer(
                        heightOfPinnedContainer,
                        bloc,
                      )
                    ],
                  ),
                );
              },
            ),
      ),
    );
  }

  Widget _buildSlivers(
    Widget stopwatch,
    bool hasExercises,
    double height,
    TrainingBloc bloc,
  ) {
    return CustomScrollView(
      slivers: [
        _buildNavBar(stopwatch, hasExercises, bloc),
        _buildTilesGrid(bloc),
        SliverToBoxAdapter(
          child: SizedBox(
            height: height,
          ),
        )
      ],
    );
  }

  Widget _buildNavBar(
    Widget stopwatch,
    bool hasExercises,
    TrainingBloc bloc,
  ) =>
      CupertinoSliverNavigationBar(
        largeTitle: Text(
          hasExercises ? 'Your Exercises' : "No exercises",
          style: TextStyle().copyWith(
            color: Theme.of(context).colorScheme.onSecondary,
          ),
        ),
        border: null,
        backgroundColor: Theme.of(context).colorScheme.surface,
        stretch: true,
        middle: stopwatch,
        trailing: BlocProvider.value(
          value: bloc,
          child: BlocBuilder<TrainingBloc, TrainingState>(
            builder: (context, state) {
              const submitWidth = 80.0; // ≈ width of the “Submit” button
              const submitHeight = 44.0; // default cupertino button height

              return Row(
                mainAxisSize: MainAxisSize.min,
                children: [
                  // CupertinoButton(
                  //   padding: EdgeInsets.zero,
                  //   onPressed: () {},
                  //   child: const Text('Random'),
                  // ),
                  // const SizedBox(width: 16),

                  SizedBox(
                    width: submitWidth,
                    height: submitHeight,
                    child: AnimatedSwitcher(
                      duration: const Duration(milliseconds: 200),
                      switchInCurve: Curves.easeIn,
                      switchOutCurve: Curves.easeOut,
                      child: CupertinoButton(
                        // “Submit”
                        key: const ValueKey(
                          'submitBtn',
                        ), // give each child a key!
                        padding: EdgeInsets.zero,
                        onPressed: hasExercises
                            ? () => bloc.add(SubmitTraining(_timerCtrl.elapsed))
                            : null,
                        child: const Text('Submit'),
                      ),
                    ),
                  ),
                ],
              );
            },
          ),
        ),
      );

  Widget _buildTilesGrid(TrainingBloc bloc) {
    // Did not use List.builder so state is saved
    // after manipulation with order
    final List<Widget> tiles = bloc.state.exercises
        .map(
          (e) => ExpandableExerciseTile(
            key: ValueKey(e.exerciseInfo.id),
            exercise: e,
            onDeleteCallback: (ex) => bloc.add(
              RemoveExercise(ex.exerciseInfo.id),
            ),
          ),
        )
        .toList();

    return SliverList(
      delegate: SliverChildListDelegate(tiles),
    );
  }

  Widget _buildPinnedContainer(double height, TrainingBloc bloc) {
    return Positioned(
      bottom: 0,
      left: 0,
      right: 0,
      child: SizedBox(
        height: height,
        child: Row(
          children: [
            Expanded(
              child: SizedBox.shrink(),
            ),
            Expanded(
              child: Center(
                child: DecoratedBox(
                  decoration: BoxDecoration(
                    color:
                        Theme.of(context).colorScheme.secondary.withAlpha(200),
                    borderRadius: BorderRadius.circular(14),
                  ),
                  child: CupertinoButton(
                    sizeStyle: CupertinoButtonSize.large,
                    onPressed: () {
                      showCupertinoModalPopup(
                        context: context,
                        builder: (context) => BlocProvider.value(
                          value: bloc,
                          child: PickingExercisePage(
                            addExerciseCallback: (info) => bloc.add(
                              AddExercise(info),
                            ),
                          ),
                        ),
                      );
                    },
                    child: Icon(
                      CupertinoIcons.add_circled,
                      size: height / 2,
                    ),
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
